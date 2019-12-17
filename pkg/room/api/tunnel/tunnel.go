package tunnel

import (
	"encoding/binary"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/Bo0km4n/arc/pkg/room/cmd/option"
	"github.com/Bo0km4n/arc/pkg/room/logger"
	"github.com/gorilla/websocket"
	"github.com/k0kubun/pp"
)

const (
	writeWait      = 10 * time.Second
	readLimit      = 30 * time.Second
	maxMessageSize = 4096 * 1000 // 4MB
	pongWait       = 60 * time.Second
	pingPeriod     = 5 * time.Second
)

var (
	T = &Tunnels{
		m:            sync.Mutex{},
		coordinators: make(map[string]*Tunnel, 4096),
		peers:        make(map[string]*Tunnel, 4096),
	}
)

type Tunnels struct {
	m            sync.Mutex
	coordinators map[string]*Tunnel
	peers        map[string]*Tunnel
}

func (t *Tunnels) StoreCoordinator(key string, conn *Tunnel) {
	t.m.Lock()
	defer t.m.Unlock()

	t.coordinators[key] = conn
}

func (t *Tunnels) LoadCoordinator(key string) (*Tunnel, bool) {
	t.m.Lock()
	defer t.m.Unlock()
	conn, ok := t.coordinators[key]
	return conn, ok
}

func (t *Tunnels) Store(key string, conn *Tunnel) {
	t.m.Lock()
	defer t.m.Unlock()
	t.peers[key] = conn
}

func (t *Tunnels) Load(key string) (*Tunnel, bool) {
	t.m.Lock()
	defer t.m.Unlock()

	conn, ok := t.peers[key]
	return conn, ok
}

type Tunnel struct {
	ttype       int // tunnel type 1=peer, 2=coordinator
	id          string
	mu          sync.Mutex
	ws          *websocket.Conn
	writeQueue  chan []byte
	Done        chan interface{}
	Err         chan error
	permissions *sync.Map
}

func NewTunnel(ttype int, ws *websocket.Conn, id string) *Tunnel {
	return &Tunnel{
		ttype:       ttype,
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
	for {
		mt, body, err := t.ws.ReadMessage()
		if mt == websocket.TextMessage {
			t.ws.SetReadDeadline(time.Now().Add(pongWait))
			continue
		}
		if mt == websocket.BinaryMessage {
			t.ws.SetReadDeadline(time.Now().Add(pongWait))
		}

		// log.Printf("receive message type: %d, ttype=%d", mt, t.ttype)
		if mt == -1 {
			// error occured
			// for debug
			if option.Opt.Debug {
				pp.Println(err)
			}
		}

		if err == io.EOF {
			logger.L.Info(fmt.Sprintf("[Peer: %s] connection closed", t.id))
			t.Done <- struct{}{}
			return nil
		}
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure, websocket.CloseNormalClosure) {
			t.Err <- err
			return err
		}
		if websocket.IsCloseError(err, 1000, 1006) {
			t.Done <- struct{}{}
			return nil
		}
		if err != nil { // catch unexpected error such as websocket.netError
			t.Err <- err
			return err
		}

		var messageType uint16
		messageType = binary.BigEndian.Uint16(body[0:2])

		switch messageType {
		case 1: // permission create
			r, err := t.CreatePermission(body[2:])
			if err != nil {
				logger.L.Error(err.Error())
				continue
			}
			t.writeQueue <- r.Raw()
		case 2: // upstream relay message
			resp, err := t.SendUpstreamRelayRequest(body[2:]) // ignore respone
			if err != nil {
				logger.L.Error(err.Error())
				continue
			}
			if resp.Status == -1 {
				t.writeQueue <- resp.Raw()
			}
		case 3: // downstream relay message
			_, err := t.SendDownstreamRelayRequest(body[2:]) // ignore response
			if err != nil {
				logger.L.Error(err.Error())
				continue
			}
		default:
			logger.L.Warn(fmt.Sprintf("Invalid message from %s", t.id))
			continue
		}
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
			if s.ttype == 2 {
				// tunnel type == with coordinator,
				// send ping message
				s.mu.Lock()
				s.ws.SetWriteDeadline(time.Now().Add(writeWait))
				if err := s.ws.WriteMessage(websocket.TextMessage, []byte("")); err != nil {
					return
				}
				s.mu.Unlock()
			}
		}
	}
}

func (t *Tunnel) OnceWriteMessage(payload []byte) {
	t.writeQueue <- payload
}

func (t *Tunnel) Close() {
	t.mu.Lock()
	t.ws.WriteMessage(websocket.CloseMessage, []byte{})
	t.ws.Close()
	t.mu.Unlock()
}

func (t *Tunnel) storePermission(peer string) {
	t.permissions.Store(peer, struct{}{})
}

func (t *Tunnel) loadPermission(peer string) bool {
	_, ok := t.permissions.Load(peer)
	return ok
}
