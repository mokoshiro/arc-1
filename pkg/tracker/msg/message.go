package msg

import "encoding/json"

type GreetMessage struct {
	PeerID    string  `json:"peer_id"`
	Mode      int     `json:"mode"`
	Addr      string  `json:"addr"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type GreetResponse struct {
	Status int `json:"status"`
}

type TrackingMessage struct {
	PeerID    string  `json:"peer_id"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type TrackingResponse struct {
	Status int `json:"status"`
}

type Message interface {
	Unmarshal([]byte) error
	Raw() []byte
}

type Response interface {
	Raw() []byte
}

func (m *GreetMessage) Raw() []byte {
	b, err := json.Marshal(m)
	if err != nil {
		return nil
	}
	return b
}

func (m *GreetMessage) Unmarshal(b []byte) error {
	return json.Unmarshal(b, m)
}

func (m *GreetResponse) Raw() []byte {
	b, err := json.Marshal(m)
	if err != nil {
		return nil
	}
	return b
}

func (m *TrackingMessage) Raw() []byte {
	b, err := json.Marshal(m)
	if err != nil {
		return nil
	}
	return b
}

func (m *TrackingMessage) Unmarshal(b []byte) error {
	return json.Unmarshal(b, m)
}

func (m *TrackingResponse) Raw() []byte {
	b, err := json.Marshal(m)
	if err != nil {
		return nil
	}
	return b
}

type LookupRequest struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Radius    float64 `json:"radius"`
	Unit      string  `json:"unit"`
	PeerID    string  `json:"peer_id"`
}

type LookupResponse struct {
	Peers []*LookupResponseRow `json:"peers"`
}

type LookupResponseRow struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	PeerID    string  `json:"peer_id"`
	Addr      string  `json:"addr"`
}

func (m *LookupRequest) Raw() []byte {
	b, err := json.Marshal(m)
	if err != nil {
		return nil
	}
	return b
}

func (m *LookupRequest) Unmarshal(b []byte) error {
	return json.Unmarshal(b, m)
}

func (m *LookupResponse) Raw() []byte {
	b, err := json.Marshal(m)
	if err != nil {
		return nil
	}
	return b
}
