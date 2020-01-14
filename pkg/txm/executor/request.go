package executor

type PreparePutPeerRequest struct {
	IdempotencyKey string  `json:"idempotency_key" binding:"required"`
	PeerID         string  `json:"peer_id" binding:"required"`
	Longitude      float64 `json:"longitude" binding:"required"`
	Latitude       float64 `json:"latitude" binding:"required"`
}

type CommitPutPeerRequest struct {
	IdempotencyKey string `json:"idempotency_key" binding:"required"`
}

type RollbackPutPeerRequest struct {
	IdempotencyKey string `json:"idempotency_key" binding:"required"`
}
