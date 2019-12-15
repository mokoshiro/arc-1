package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Bo0km4n/arc/pkg/room/api/middleware"
	"github.com/Bo0km4n/arc/pkg/room/api/tunnel"
	"github.com/Bo0km4n/arc/pkg/room/cmd/option"
	"github.com/Bo0km4n/arc/pkg/room/logger"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{}
)

func Accept(w http.ResponseWriter, r *http.Request) {
	logger.L.Info("Accepted Connection")
	peerID := r.Header.Get("X-ARC-PEER-ID")
	credential := r.Header.Get("X-ARC-PEER-CREDENTIAL")

	if err := middleware.ValidateCredential(peerID, credential); err != nil {
		logger.L.Error(err.Error())
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.L.Error(err.Error())
		return
	}
	defer ws.Close()
	t := tunnel.NewTunnel(ws)
	tunnel.T.Store(peerID, t)

	go t.ReadPump()
	go t.WritePump()

	select {
	case _, ok := <-t.Done:
		if !ok {
			t.Close()
		}
		logger.L.Info("Closing websocket")
		return
	}
}

func Run(ctx context.Context) error {
	http.HandleFunc("/bind", func(w http.ResponseWriter, r *http.Request) {
		Accept(w, r)
	})
	logger.L.Info(fmt.Sprintf("Listening Room Server on [:%d]", option.Opt.Port))
	return http.ListenAndServe(fmt.Sprintf(":%d", option.Opt.Port), nil)
}
