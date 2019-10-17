package repository

import (
	"github.com/Bo0km4n/arc/pkg/metadata/cmd/option"
	"github.com/Bo0km4n/arc/pkg/metadata/infra/db"
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

func (rr *registerKVSRepository) Register(peerID string, addr string) error {
	conn := rr.redisPool.Get()
	_, err := conn.Do("SETEX", peerID, option.Opt.RedisKeyExpire, addr)
	return err
}
