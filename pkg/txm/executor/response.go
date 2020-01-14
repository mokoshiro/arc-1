package executor

type PreparePutPeerResponse struct {
	IdempotencyKey string
}

type CommitPutPeerResponse struct {
	IdempotencyKey string
}

type RollbackPutPeerResponse struct {
	IdempotencyKey string
}
