package model

// RegisterRequest POST /api/register
type RegisterRequest struct {
	GlobalIPAddr string    `json:"global_ip_addr"`
	Port         string    `json:"port"`
	ID           string    `json:"id"`
	Location     *Location `json:"location"`
}

type RegisterResponse struct {
}
