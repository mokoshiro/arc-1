package model

// Location includes latitude and longitude
type Location struct {
	Latitude  float64 `json:"latitude" form:"latitude"`
	Longitude float64 `json:"longitude" form:"longitude"`
}
