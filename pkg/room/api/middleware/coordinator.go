package middleware

import (
	"log"
	"net/http"
	"net/url"

	"github.com/Bo0km4n/arc/pkg/room/infra/db"
	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/websocket"
)

func DialOtherCoordinator(selfAddr, addr string) (*websocket.Conn, error) {
	header := http.Header{}
	header.Add("X-ARC-COORDINATOR-ID", selfAddr)
	u := url.URL{Scheme: "ws", Host: addr, Path: "/coordinator"}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), header)
	if err != nil {
		log.Fatal(err)
	}

	return c, err
}

func GetRemoteCoordinatorAddressByPeer(id string) (string, error) {
	c := db.RedisPool.Get()
	defer c.Close()

	return redis.String(c.Do("GET", id))
}
