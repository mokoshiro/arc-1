package model

// RegisterRequest POST /api/register
type RegisterRequest struct {
	GlobalIPAddr string    `json:"global_ip_addr" binding:"required"`
	Port         string    `json:"port" binding:"required"`
	ID           string    `json:"id" binding:"required"`
	Location     *Location `json:"location" binding:"required"`
}

type RegisterResponse struct {
}
