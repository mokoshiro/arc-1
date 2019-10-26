package repository

import (
	"context"

	"github.com/Bo0km4n/arc/pkg/metadata/cmd/option"
	"github.com/Bo0km4n/arc/pkg/metadata/infra/db"
	"github.com/garyburd/redigo/redis"
)

type MemberRepository interface {
	Register(peerID, addr string) error
	GetMember(ctx context.Context, peerIDs []string) ([]string, error)
	Delete(ctx context.Context, peerID string) error
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

func (mr *memberKVSRepository) GetMember(ctx context.Context, peerIDs []string) ([]string, error) {
	conn := mr.redisPool.Get()
	defer conn.Close()
	if len(peerIDs) == 1 {
		res, err := redis.String(conn.Do("GET", peerIDs[0]))
		if err != nil {
			return []string{}, err
		}
		return []string{res}, nil
	}
	args := make([]interface{}, len(peerIDs))
	for i := range peerIDs {
		args[i] = peerIDs[i]
	}
	return redis.Strings(conn.Do("MGET", args...))
}

func (mr *memberKVSRepository) Delete(ctx context.Context, peerID string) error {
	conn := mr.redisPool.Get()
	defer conn.Close()
	if _, err := conn.Do("DEL", peerID); err != nil {
		return err
	}
	return nil
}
