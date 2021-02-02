package executor

// Insert Peer

type PreparePutPeerRequest struct {
	PeerID     string  `json:"peer_id" binding:"required"`
	Addr       string  `json:"addr" binding:"required"`
	Credential string  `json:"credential" binding:"required"`
	Longitude  float64 `json:"longitude" binding:"required"`
	Latitude   float64 `json:"latitude" binding:"required"`
}

// Update Peer

type UpdatePeerRequest struct {
	PeerID     string  `json:"peer_id" binding:"required"`
	Addr       string  `json:"addr" binding:"required"`
	Credential string  `json:"credential" binding:"required"`
	Longitude  float64 `json:"longitude" binding:"required"`
	Latitude   float64 `json:"latitude" binding:"required"`
}

type CommitUpdatePeerRequest struct {
	IdempotencyKey string `json:"idempotency_key" binding:"required"`
}

type RollbackUpdatePeerRequest struct {
	IdempotencyKey string `json:"idempotency_key" binding:"required"`
}

// Delete Peer
type DeletePeerRequest struct {
	PeerID string `json:"peer_id"`
}

type DeletePeerResponse struct{}

// Util

type ResourceUsage struct {
}

type LookupRequest struct {
	Radius    float64 `json:"radius" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required"`
}


