package message

import (
	"encoding/binary"
	"encoding/json"
)

type UpstreamRelayRequest struct {
	DestinationLength uint32
	Destinations      *Destinations
	Payload           []byte
}

type Destinations struct {
	Peers []string `peers`
}

func ParseUpstreamRelayRequest(body []byte) (*UpstreamRelayRequest, error) {
	length := binary.BigEndian.Uint32(body[0:4])
	dests := &Destinations{}
	if err := json.Unmarshal(body[4:4+length], dests); err != nil {
		return nil, err
	}
	return &UpstreamRelayRequest{
		DestinationLength: length,
		Destinations:      dests,
		Payload:           body[4+length:],
	}, nil
}

type UpstreamRelayResponse struct {
	Status int `json:"status"`
}

func (ur *UpstreamRelayResponse) Raw() []byte {
	b, _ := json.Marshal(ur)
	return b
}
