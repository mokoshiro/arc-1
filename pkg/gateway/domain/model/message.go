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

type GetMemberByRadiusRequest struct {
	Location  *Location `form:"location" binding:"required"`
	Radius    float64   `form:"radius" binding:"required"`
	WithCoord bool      `form:"with_coord" binding:"required"`
	Unit      string    `form:"unit" binding:"required"`
}

type GetMemberByRadiusResponse struct {
	Members []*Member `json:"members"`
}

type Member struct {
	Location *Location `json:"location"`
	Addr     string    `json:"string"`
	ID       string    `json:"id"`
}
