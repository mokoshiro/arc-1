package socket

import (
	"encoding/binary"
	"encoding/json"
	"errors"
)

type Message interface {
	Unmarshal([]byte) error
	Raw() []byte
}

type Response interface {
	Raw() []byte
}

type greetMessage struct {
	peerID    string  `json:"peer_id"`
	mode      int     `json:"mode"`
	addr      string  `json:"addr"`
	longitude float64 `json:"longitude"`
	latitude  float64 `json:"latitude"`
}

type greetResponse struct {
	Status int `json:"status"`
}

func (m *greetMessage) Raw() []byte {
	b, err := json.Marshal(m)
	if err != nil {
		return nil
	}
	return b
}

func (m *greetMessage) Unmarshal(b []byte) error {
	return json.Unmarshal(b, m)
}

func (m *greetResponse) Raw() []byte {
	b, err := json.Marshal(m)
	if err != nil {
		return nil
	}
	return b
}

func handleMessage(body []byte) (Response, error) {
	var messageType uint16
	messageType = binary.BigEndian.Uint16(body[0:2])
	switch messageType {
	case 0:
		return nil, errors.New("Unimplemented message type: 0")
	case 1:
		gm := &greetMessage{}
		if err := gm.Unmarshal(body[4:]); err != nil {
			return nil, err
		}
		return &greetResponse{Status: 1}, nil
	}
	return nil, nil
}
