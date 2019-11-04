package socket

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type Message interface {
	Marshal([]byte) error
	Raw() []byte
}

type Response interface {
	Raw() []byte
}

type greetMessage struct {
	peerID    string
	mode      int
	addr      string
	longitude float64
	latitude  float64
}

type greetResponse struct {
	Status int
}

func (m *greetMessage) Raw() []byte {
	buf := &bytes.Buffer{}
	if err := binary.Write(buf, binary.BigEndian, &m); err != nil {
		return nil
	}
	return buf.Bytes()
}

func (m *greetMessage) Marshal(b []byte) error {
	reader := bytes.NewReader(b)
	return binary.Read(reader, binary.BigEndian, m)
}

func (m *greetResponse) Raw() []byte {
	buf := &bytes.Buffer{}
	if err := binary.Write(buf, binary.BigEndian, &m); err != nil {
		return nil
	}
	return buf.Bytes()
}

func handleMessage(body []byte) (Response, error) {
	var messageType uint16
	binary.BigEndian.PutUint16(body[0:2], messageType)
	switch messageType {
	case 0:
		return nil, errors.New("Unimplemented message type: 0")
	case 1:
		gm := &greetMessage{}
		if err := gm.Marshal(body[4:]); err != nil {
			return nil, err
		}
		return &greetResponse{Status: 1}, nil
	}
	return nil, nil
}
