package middleware

import (
	"fmt"

	"github.com/Bo0km4n/arc/pkg/room/cmd/option"
	"github.com/Bo0km4n/arc/pkg/room/infra/db"
)

func SetTunnel(peerID string, roomAddr string) error {
	redisConn := db.RedisPool.Get()
	defer redisConn.Close()

	_, err := redisConn.Do("SETEX", peerID, option.Opt.RedisKeyExpire, fmt.Sprintf("%s:%d", roomAddr, option.Opt.Port))
	return err
}

func RemoveTunnel(peerID string) {
	redisConn := db.RedisPool.Get()
	defer redisConn.Close()
	redisConn.Do("DEL", peerID)
}
