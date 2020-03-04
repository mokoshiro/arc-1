package driver

type RegisterResponse struct {
	Message string
}

type LookupResponse struct {
	Peers []struct {
		PeerID    string  `json:"peer_id" binding:"required"`
		Longitude float64 `json:"longitude" binding:"required"`
		Latitude  float64 `json:"latitude" binding:"required"`
	} `json:"peers"`
}
