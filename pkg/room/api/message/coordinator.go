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
	head := []byte{0x00, 0x04}
	b, _ := json.Marshal(ur)
	head = append(head, b...)
	return head
}

type DownstreamRelayRequest struct {
	SourceLength      uint32
	Source            string
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
	sourceLength := binary.BigEndian.Uint32(body[0:4])
	source := string(body[4 : 4+sourceLength])
	destLength := binary.BigEndian.Uint32(body[4+sourceLength : 8+sourceLength])
	dest := string(body[8+sourceLength : 8+sourceLength+destLength])
	return &DownstreamRelayRequest{
		SourceLength:      sourceLength,
		Source:            source,
		DestinationLength: destLength,
		Destination:       dest,
		Payload:           body[8+sourceLength+destLength:],
	}, nil
}

func NewDownstreamRelayRequestRaw(source string, dest string, payload []byte) []byte {
	head := []byte{0x00, 0x03} // downstream request type
	sourceLen := uint32(len(source))
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, sourceLen)
	head = append(head, bytes...)
	head = append(head, []byte(source)...)
	destLen := uint32(len(dest))
	binary.BigEndian.PutUint32(bytes, destLen)
	head = append(head, bytes...)
	head = append(head, []byte(dest)...)
	head = append(head, payload...)
	return head
}
