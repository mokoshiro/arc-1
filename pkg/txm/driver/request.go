package driver

type RegisterRequest struct {
	PeerID    string  `json:"peer_id" binding:"required"`
	Addr      string  `json:"addr" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required"`
}

type UpdatePeerLocationRequest struct {
	PeerID    string  `json:"peer_id" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required"`
}

type LookupRequest struct {
	Radius    float64 `json:"radius" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required"`
}
