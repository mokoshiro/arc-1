package repository

import (
	"github.com/Bo0km4n/arc/pkg/tracker/infra/db"
	"github.com/garyburd/redigo/redis"
)

type RegisterRepository interface {
	Register(h3Hash, peerID string, longitude, latitude float64) error
}

type registerKVSRepository struct {
	redisPool *redis.Pool
}

type registerRDBRepository struct {
	// *sql.DB
}

func newRegisterKVSRepository(redisPool *redis.Pool) RegisterRepository {
	return &registerKVSRepository{
		redisPool: redisPool,
	}
}

func NewRegisterRepository(dbType int) RegisterRepository {
	switch dbType {
	case db.DB_REDIS:
		return newRegisterKVSRepository(db.RedisPool)
	default:
		return nil
	}
}

func (rr *registerKVSRepository) Register(h3Hash, peerID string, latitude, longitude float64) error {
	conn := rr.redisPool.Get()
	_, err := conn.Do("GEOADD", h3Hash, longitude, longitude, peerID)
	return err
}
