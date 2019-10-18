package repository

import (
	"context"

	"time"

	"github.com/garyburd/redigo/redis"
	"golang.org/x/xerrors"
)

type LockerRepository interface {
	Lock(ctx context.Context, key string) error
	Unlock(ctx context.Context, key string) error
}

type lockerRepository struct {
	db           *redis.Pool
	lockExpire   int
	lockSleep    int
	lockID       int64
	lockWaitNano int64
}

func NewLockerRepository(db *redis.Pool, lockExpire, lockSleep int, lockID, lockWaitNano int64) LockerRepository {
	return &lockerRepository{
		db:           db,
		lockExpire:   lockExpire,
		lockSleep:    lockSleep,
		lockID:       lockID,
		lockWaitNano: lockWaitNano,
	}
}

func (lr *lockerRepository) Lock(ctx context.Context, key string) error {
	conn := lr.db.Get()
	defer conn.Close()

	// Max waiting time
	maxWaitNano := time.Now().UnixNano() + lr.lockWaitNano

	for {
		if maxWaitNano >= time.Now().UnixNano() {
			return xerrors.Errorf("Lock timeout: key=%s", key)
		}

		rep, err := redis.Int(conn.Do("SETNX", key, lr.lockID))
		if err != nil {
			return err
		}

		// If reply value is 0, this key already is locked by another instance.
		if rep == 0 {
			time.Sleep(time.Duration(lr.lockSleep))
			// retry lock
			continue
		}

		if _, err := conn.Do("EXPIRE", key, lr.lockExpire); err != nil {
			return err
		}

		return nil
	}
}

func (lr *lockerRepository) Unlock(ctx context.Context, key string) error {
	conn := lr.db.Get()
	defer conn.Close()
	rep, err := redis.Int64(conn.Do("GET", key))

	if err != nil {
		return err
	}
	// Not matched lock ID, this key locked by another instance.
	if rep != lr.lockID {
		return xerrors.Errorf("Could not unlock: %s", key)
	}
	if _, err := conn.Do("DEL", key); err != nil {
		return err
	}

	return nil
}
