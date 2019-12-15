package tunnel

import (
	"github.com/Bo0km4n/arc/pkg/room/api/message"
	"github.com/k0kubun/pp"
)

func (t *Tunnel) SendUpstreamRelayRequest(body []byte) (*message.UpstreamRelayResponse, error) {
	req, err := message.ParseUpstreamRelayRequest(body)
	if err != nil {
		return nil, err
	}
	pp.Println(req)
	return &message.UpstreamRelayResponse{
		Status: 1,
	}, nil
}
