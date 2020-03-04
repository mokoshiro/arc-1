package schema

import (
	"time"
)

type Peer struct {
	PeerID     string  `json:"peer_id"`
	Addr       string  `json:"addr"`
	Credential string  `json:"credential"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	CreatedAt  time.Time
	UpdatedAt  *time.Time
}
