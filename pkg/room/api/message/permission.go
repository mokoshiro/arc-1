package message

import (
	"encoding/json"
)

type PermissionRequest struct {
	Peers []string `json:"peers"`
}

type PermissionResponse struct {
	Status int `json:"status"`
}

func ParsePermissionRequest(b []byte) (*PermissionRequest, error) {
	req := &PermissionRequest{}
	if err := json.Unmarshal(b, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (pr *PermissionResponse) Raw() []byte {
	b, _ := json.Marshal(pr)
	return b
}
