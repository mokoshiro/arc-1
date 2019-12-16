package tunnel

import (
	"encoding/binary"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/Bo0km4n/arc/pkg/room/api/message"
	"github.com/Bo0km4n/arc/pkg/room/logger"
	"github.com/gorilla/websocket"
	"github.com/k0kubun/pp"
)

const (
	writeWait      = 10 * time.Second
	readLimit      = 60 * time.Second
	maxMessageSize = 512
	pongWait       = 60 * time.Second
	pingPeriod     = 10 * time.Second
)

var (
	T = &Tunnels{
		m:     sync.Mutex{},
		peers: make(map[string]*Tunnel, 4096),
	}
)

type Tunnels struct {
	m            sync.Mutex
	coordinators map[string]*websocket.Conn
	peers        map[string]*Tunnel
}

func (t *Tunnels) StoreCoordinator(key string, conn *websocket.Conn) {
	t.m.Lock()
	defer t.m.Unlock()
	t.coordinators[key] = conn
}

func (t *Tunnels) LoadCoordinator(key string) (*websocket.Conn, error) {
	t.m.Lock()
	defer t.m.Unlock()
	conn, ok := t.coordinators[key]
	if !ok {
		return nil, fmt.Errorf("Coordinator: %s is not found", key)
	}
	return conn, nil
}

func (t *Tunnels) Store(key string, conn *Tunnel) {
	t.m.Lock()
	defer t.m.Unlock()
	t.peers[key] = conn
}

func (t *Tunnels) Load(key string) (*Tunnel, error) {
	t.m.Lock()
	defer t.m.Unlock()

	conn, ok := t.peers[key]
	if !ok {
		return nil, fmt.Errorf("Peer ID: %s is not found", key)
	}
	return conn, nil
}

type Tunnel struct {
	id          string
	mu          sync.Mutex
	ws          *websocket.Conn
	writeQueue  chan []byte
	Done        chan interface{}
	Err         chan error
	permissions *sync.Map
}

func NewTunnel(ws *websocket.Conn, id string) *Tunnel {
	return &Tunnel{
		id:          id,
		ws:          ws,
		writeQueue:  make(chan []byte, 1),
		Done:        make(chan interface{}, 1),
		Err:         make(chan error, 1),
		permissions: new(sync.Map),
	}
}

func (t *Tunnel) ReadPump() error {
	t.ws.SetReadLimit(maxMessageSize)
	t.ws.SetReadDeadline(time.Now().Add(readLimit))
	t.ws.SetPongHandler(func(string) error { t.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, body, err := t.ws.ReadMessage()
		if err == io.EOF {
			logger.L.Info(fmt.Sprintf("[Peer: %s] connection closed", t.id))
			t.Done <- struct{}{}
			return nil
		}
		if err != nil && websocket.IsUnexpectedCloseError(err, 1000) {
			logger.L.Info(err.Error())
			t.Err <- err
			return err
		}
		if websocket.IsCloseError(err, 1000) {
			t.Done <- struct{}{}
			return nil
		}

		var messageType uint16
		messageType = binary.BigEndian.Uint16(body[0:2])

		var resp message.Response
		pp.Println(messageType)
		switch messageType {
		case 1: // permission create
			r, err := t.CreatePermission(body[2:])
			if err != nil {
				logger.L.Error(err.Error())
				continue
			}
			resp = r
		case 2: // upstream relay message
			r, err := t.SendUpstreamRelayRequest(body[2:])
			if err != nil {
				logger.L.Error(err.Error())
				continue
			}
			resp = r
		default:
			logger.L.Warn(fmt.Sprintf("Invalid message from %s", t.id))
			continue
		}

		t.writeQueue <- resp.Raw()
	}
}

func (s *Tunnel) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		s.ws.Close()
	}()

	for {
		select {
		case message, ok := <-s.writeQueue:
			s.mu.Lock()
			if !ok {
				s.ws.WriteMessage(websocket.CloseMessage, []byte{})
				logger.L.Error("Failed get message from write queue")
				return
			}
			if err := s.ws.WriteMessage(websocket.BinaryMessage, message); err != nil {
				logger.L.Error(err.Error())
				return
			}
			s.mu.Unlock()

		case <-ticker.C:
			s.mu.Lock()
			s.ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := s.ws.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
			s.mu.Unlock()
		}
	}
}

func (t *Tunnel) Close() {
	t.mu.Lock()
	t.ws.WriteMessage(websocket.CloseMessage, []byte{})
	t.mu.Unlock()
}

func (t *Tunnel) storePermission(peer string) {
	t.permissions.Store(peer, struct{}{})
}

func (t *Tunnel) loadPermission(peer string) bool {
	_, ok := t.permissions.Load(peer)
	return ok
}
