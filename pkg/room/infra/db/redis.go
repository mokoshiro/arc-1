package db

import (
	"time"

	"github.com/Bo0km4n/arc/pkg/room/cmd/option"
	"github.com/garyburd/redigo/redis"
)

var (
	RedisPool *redis.Pool
)

func InitRedisPool() {
	RedisPool = &redis.Pool{
		Wait:        true,
		MaxIdle:     option.Opt.RedisMaxIdle,
		MaxActive:   option.Opt.RedisActive,
		IdleTimeout: time.Duration(option.Opt.RedisIdleTimeout) * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", option.Opt.RedisHost) },
	}
}
