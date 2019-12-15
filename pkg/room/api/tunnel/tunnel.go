package tunnel

import (
	"fmt"
	"io"
	"sync"
	"time"

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
	m     sync.Mutex
	peers map[string]*Tunnel
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
		return nil, fmt.Errorf("Peer ID: [%s] is not found", key)
	}
	return conn, nil
}

type Tunnel struct {
	id          string
	ws          *websocket.Conn
	writeQueue  chan []byte
	Done        chan interface{}
	Err         chan error
	permissions []string
}

func NewTunnel(ws *websocket.Conn, id string) *Tunnel {
	return &Tunnel{
		id:          id,
		ws:          ws,
		writeQueue:  make(chan []byte, 1),
		Done:        make(chan interface{}, 1),
		Err:         make(chan error, 1),
		permissions: make([]string, 1024),
	}
}

func (s *Tunnel) ReadPump() error {
	s.ws.SetReadLimit(maxMessageSize)
	s.ws.SetReadDeadline(time.Now().Add(readLimit))
	s.ws.SetPongHandler(func(string) error { s.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, body, err := s.ws.ReadMessage()
		if err == io.EOF {
			logger.L.Info(fmt.Sprintf("[Peer: %s] connection closed", s.id))
			s.Done <- struct{}{}
			return nil
		}
		if err != nil {
			logger.L.Info(err.Error())
			s.Err <- err
			return err
		}
		pp.Println(body)
		// resp, err := handleMessage(body)
		// if err != nil {
		// 	logger.L.Error(err.Error())
		// 	continue
		// }
		// s.writeQueue <- resp.Raw()
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
			if !ok {
				s.ws.WriteMessage(websocket.CloseMessage, []byte{})
				logger.L.Error("Failed get message from write queue")
				return
			}
			if err := s.ws.WriteMessage(websocket.BinaryMessage, message); err != nil {
				logger.L.Error(err.Error())
				return
			}
		case <-ticker.C:
			s.ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := s.ws.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (t *Tunnel) Close() {
	t.ws.WriteMessage(websocket.CloseMessage, []byte{})
}
