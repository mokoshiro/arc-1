package message

import (
	"encoding/binary"
	"encoding/json"

	"github.com/k0kubun/pp"
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
	head := []byte{0x00, 0x04}
	b, _ := json.Marshal(ur)
	head = append(head, b...)
	return head
}

type DownstreamRelayRequest struct {
	DestinationLength uint32
	Destination       string
	Payload           []byte
}

type DownstreamRelayResponse struct {
	Status int `json:"status"`
}

func (dr *DownstreamRelayResponse) Raw() []byte {
	b, _ := json.Marshal(dr)
	return b
}

func ParseDownstreamRelayRequest(body []byte) (*DownstreamRelayRequest, error) {
	pp.Println(string(body))
	length := binary.BigEndian.Uint32(body[0:4])
	dest := string(body[4 : 4+length])

	return &DownstreamRelayRequest{
		DestinationLength: length,
		Destination:       dest,
		Payload:           body[4+length:],
	}, nil
}

func NewDownstreamRelayRequestRaw(dest string, payload []byte) []byte {
	head := []byte{0x00, 0x03} // downstream request type
	destLen := uint32(len(dest))
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, destLen)
	head = append(head, bytes...)
	head = append(head, []byte(dest)...)
	head = append(head, payload...)
	return head
}
