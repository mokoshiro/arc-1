package driver

type PutRequest struct {
	PeerID    string  `json:"peer_id" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required"`
}

type LookupRequest struct {
	Radius    float64 `json:"radius" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required"`
}
