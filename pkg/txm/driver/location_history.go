package driver

import "github.com/garyburd/redigo/redis"

type locationHistory struct {
	kvs *redis.Pool
}

func newLocationHistory(kvs *redis.Pool) *locationHistory {
	return &locationHistory{
		kvs: kvs,
	}
}

func (lh *locationHistory) Put(peerID string, executorAddr string) error {
	rc := lh.kvs.Get()
	_, err := rc.Do("SET", peerID, executorAddr)
	return err
}

func (lh *locationHistory) Get(peerID string) (string, error) {
	rc := lh.kvs.Get()
	return redis.String(rc.Do("GET", peerID))
}
