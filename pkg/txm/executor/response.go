package executor

type PreparePutPeerResponse struct {
}

type UpdatePeerLocationResponse struct {
}

type ResourceUsageResponse struct {
	MemUsedPercent float64 `json:"mem_used_percent"`
	CpuUsedPercent float64 `json:"cpu_used_percent"`
}
