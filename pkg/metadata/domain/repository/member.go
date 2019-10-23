package repository

import (
	"github.com/Bo0km4n/arc/pkg/metadata/cmd/option"
	"github.com/Bo0km4n/arc/pkg/metadata/infra/db"
	"github.com/garyburd/redigo/redis"
)

type MemberRepository interface {
	Register(peerID, addr string) error
}

type memberKVSRepository struct {
	redisPool *redis.Pool
}

type memberRDBRepository struct {
	// *sql.DB
}

func newMemberKVSRepository(redisPool *redis.Pool) MemberRepository {
	return &memberKVSRepository{
		redisPool: redisPool,
	}
}

func NewMemberRepository(dbType int) MemberRepository {
	switch dbType {
	case db.DB_REDIS:
		return newMemberKVSRepository(db.RedisPool)
	default:
		return nil
	}
}

func (rr *memberKVSRepository) Register(peerID string, addr string) error {
	conn := rr.redisPool.Get()
	defer conn.Close()
	_, err := conn.Do("SETEX", peerID, option.Opt.RedisKeyExpire, addr)
	return err
}
