package repository

import (
	"github.com/Bo0km4n/arc/pkg/metadata/cmd/option"
	"github.com/garyburd/redigo/redis"
)

type RegisterRepository interface {
	Register(peerID, addr string) error
}

type registerKVSRepository struct {
	redisPool *redis.Pool
}

type registerRDBRepository struct {
	// *sql.DB
}

func NewRegisterKVSRepository(redisPool *redis.Pool) RegisterRepository {
	return &registerKVSRepository{
		redisPool: redisPool,
	}
}

func (rr *registerKVSRepository) Register(peerID string, addr string) error {
	conn := rr.redisPool.Get()
	_, err := conn.Do("SETEX", peerID, option.Opt.RedisKeyExpire, addr)
	return err
}
