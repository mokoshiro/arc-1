package tunnel

import (
	"fmt"
	"github.com/Bo0km4n/arc/pkg/room/api/message"
	"github.com/Bo0km4n/arc/pkg/room/api/middleware"
	"github.com/Bo0km4n/arc/pkg/room/cmd/option"
	"github.com/Bo0km4n/arc/pkg/room/logger"
	"github.com/k0kubun/pp"
)

func (t *Tunnel) SendUpstreamRelayRequest(body []byte) (*message.UpstreamRelayResponse, error) {
	req, err := message.ParseUpstreamRelayRequest(body)
	if err != nil {
		return nil, err
	}
	pp.Println(req)

	for _, dest := range req.Destinations.Peers {
		if err := t.relayUpstreamRequest(dest, req.Payload); err != nil {
			logger.L.Warn(err.Error())
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
		remoteTunnel = NewTunnel(ws, remoteCoordinatorAddr)
		T.StoreCoordinator(remoteCoordinatorAddr, t)
	}
	remoteTunnel.OnceWriteMessage(payload)
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
	leftTunnel.OnceWriteMessage(req.Payload)
	return &message.DownstreamRelayResponse{
		Status: 1,
	}, nil
}
