package socket

import (
	"net/http"

	"time"

	"github.com/Bo0km4n/arc/pkg/tracker/logger"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var conns = map[string]*Sock{}

const (
	writeWait      = 10 * time.Second
	readLimit      = 60 * time.Second
	maxMessageSize = 512
	pongWait       = 60 * time.Second
	pingPeriod     = 10 * time.Second
)

type Sock struct {
	ws         *websocket.Conn
	writeQueue chan []byte
	done       chan interface{}
}

func Serve(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.L.Error(err.Error())
		return
	}
	defer ws.Close()
	sock := newSocket(ws)

	go sock.readPump()
	go sock.writePump()

	select {
	case _, ok := <-sock.done:
		if !ok {
			sock.ws.WriteMessage(websocket.CloseMessage, []byte{})
		}
		logger.L.Info("Closing websocket")
		return
	}
}

func newSocket(ws *websocket.Conn) *Sock {
	return &Sock{
		ws:         ws,
		writeQueue: make(chan []byte, 1),
		done:       make(chan interface{}, 1),
	}
}

func (s *Sock) readPump() error {
	s.ws.SetReadLimit(maxMessageSize)
	s.ws.SetReadDeadline(time.Now().Add(readLimit))
	s.ws.SetPongHandler(func(string) error { s.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, body, err := s.ws.ReadMessage()
		if err != nil {
			logger.L.Info(err.Error())
			return err
		}
		resp, err := handleMessage(body)
		if err != nil {
			logger.L.Error(err.Error())
			continue
		}
		s.writeQueue <- resp.Raw()
	}
}

func (s *Sock) writePump() {
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
