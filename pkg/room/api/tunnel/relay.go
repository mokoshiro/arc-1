package tunnel

import (
	"github.com/Bo0km4n/arc/pkg/room/api/message"
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
	return nil
}
