package tunnel

import "github.com/Bo0km4n/arc/pkg/room/api/message"

func (t *Tunnel) CreatePermission(b []byte) (*message.PermissionResponse, error) {
	req, err := message.ParsePermissionRequest(b)
	if err != nil {
		return nil, err
	}

	for _, v := range req.Peers {
		t.storePermission(v)
	}
	return &message.PermissionResponse{
		Type:   "Permission Create",
		Status: 1,
	}, nil
}

func (t *Tunnel) IsExistPermission(key string) bool {
	return t.loadPermission(key)
}
