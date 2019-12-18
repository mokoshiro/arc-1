package tunnel

import (
	"fmt"

	"github.com/Bo0km4n/arc/pkg/room/api/message"
	"github.com/Bo0km4n/arc/pkg/room/api/middleware"
	"github.com/Bo0km4n/arc/pkg/room/cmd/option"
	"github.com/Bo0km4n/arc/pkg/room/logger"
)

func (t *Tunnel) SendUpstreamRelayRequest(body []byte) (*message.UpstreamRelayResponse, error) {
	req, err := message.ParseUpstreamRelayRequest(body)
	if err != nil {
		return nil, err
	}

	for _, dest := range req.Destinations.Peers {
		if err := t.relayUpstreamRequest(dest, req.Payload); err != nil {
			logger.L.Warn(err.Error())
			return &message.UpstreamRelayResponse{
				Status: -1,
			}, nil
		}
	}
	return &message.UpstreamRelayResponse{
		Status: 1,
	}, nil
}

func (t *Tunnel) relayUpstreamRequest(dest string, payload []byte) error {
	// get remote coordinator address from redis
	remoteCoordinatorAddr, err := middleware.GetRemoteCoordinatorAddressByPeer(dest)
	if err != nil {
		return err
	}
	remoteTunnel, ok := T.LoadCoordinator(remoteCoordinatorAddr)
	if !ok {
		addr := fmt.Sprintf("%s:%d", option.Opt.GlobalAddress, option.Opt.Port)
		ws, err := middleware.DialOtherCoordinator(addr, remoteCoordinatorAddr)
		if err != nil {
			return err
		}
		remoteTunnel = NewTunnel(2, ws, remoteCoordinatorAddr)
		go remoteTunnel.ReadPump()
		go remoteTunnel.WritePump()
		T.StoreCoordinator(remoteCoordinatorAddr, remoteTunnel)
	}
	req := message.NewDownstreamRelayRequestRaw(t.id, dest, payload)
	remoteTunnel.OnceWriteMessage(req)
	return nil
}

func (t *Tunnel) SendDownstreamRelayRequest(body []byte) (*message.DownstreamRelayResponse, error) {
	req, err := message.ParseDownstreamRelayRequest(body)
	if err != nil {
		return nil, err
	}
	leftTunnel, ok := T.Load(req.Destination)
	if !ok {
		return nil, fmt.Errorf("Not found peer connection: id=%s", req.Destination)
	}
	if !leftTunnel.IsExistPermission(req.Source) {
		return nil, fmt.Errorf("Access denied from id=%s", req.Source)
	}
	leftTunnel.OnceWriteMessage(req.Payload)
	return &message.DownstreamRelayResponse{
		Status: 1,
	}, nil
}
