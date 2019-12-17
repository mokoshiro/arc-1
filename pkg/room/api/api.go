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

func AcceptPeer(w http.ResponseWriter, r *http.Request) {
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
	// Build ws connection tunnel object
	t := tunnel.NewTunnel(1, ws, peerID)
	// Set connection information to Relay DNS
	if err := middleware.SetTunnel(peerID); err != nil {
		logger.L.Error(err.Error())
		return
	}
	// Store tunnel object to global dictionary
	tunnel.T.Store(peerID, t)

	go t.ReadPump()
	go t.WritePump()

	select {
	case _, ok := <-t.Done:
		if !ok {
			t.Close()
		}
		logger.L.Info(fmt.Sprintf("Closing websocket: peer_id=%s", peerID))
		t.Close()
		middleware.RemoveTunnel(peerID)
		return
	case err, _ := <-t.Err:
		e := fmt.Sprintf("Unexpected error is happened in websockets: %s, peer_id=%s", err.Error(), peerID)
		logger.L.Error(e)
		t.Close()
		middleware.RemoveTunnel(peerID)
		return
	}
}

func AcceptCoordinator(w http.ResponseWriter, r *http.Request) {
	remoteCoordinatorAddr := r.Header.Get("X-ARC-COORDINATOR-ID")
	ws, err := upgrader.Upgrade(w, r, nil)
	logger.L.Info(fmt.Sprintf("Accepted Coordinator connection: addr=%s", remoteCoordinatorAddr))
	if err != nil {
		logger.L.Error(err.Error())
		return
	}
	defer ws.Close()

	coordinatorTunnel := tunnel.NewTunnel(2, ws, remoteCoordinatorAddr)
	tunnel.T.StoreCoordinator(remoteCoordinatorAddr, coordinatorTunnel)

	go coordinatorTunnel.ReadPump()
	go coordinatorTunnel.WritePump()

	select {
	case _, ok := <-coordinatorTunnel.Done:
		if !ok {
			coordinatorTunnel.Close()
		}
		logger.L.Info(fmt.Sprintf("Closing websocket: coordinator_address=%s", remoteCoordinatorAddr))
		coordinatorTunnel.Close()
		return
	case err, _ := <-coordinatorTunnel.Err:
		e := fmt.Sprintf("Unexpected error is happened in websockets: %s, peer_id=%s", err.Error(), remoteCoordinatorAddr)
		logger.L.Error(e)
		coordinatorTunnel.Close()
		return
	}
}

func Run(ctx context.Context) error {
	http.HandleFunc("/bind", func(w http.ResponseWriter, r *http.Request) {
		AcceptPeer(w, r)
	})
	http.HandleFunc("/coordinator", func(w http.ResponseWriter, r *http.Request) {
		AcceptCoordinator(w, r)
	})
	logger.L.Info(fmt.Sprintf("Listening Room Server on [:%d]", option.Opt.Port))
	return http.ListenAndServe(fmt.Sprintf(":%d", option.Opt.Port), nil)
}
