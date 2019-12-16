package middleware

import (
	"log"
	"net/http"
	"net/url"

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
