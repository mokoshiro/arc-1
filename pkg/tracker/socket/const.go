package socket

import (
	"fmt"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

var (
	upgrader  = websocket.Upgrader{}
	conns     = map[string]*Sock{}
	SOCK_ADDR = fmt.Sprintf("%s:%s", os.Getenv("POD_IP"), "8000")
)

const (
	writeWait      = 10 * time.Second
	readLimit      = 60 * time.Second
	maxMessageSize = 512
	pongWait       = 60 * time.Second
	pingPeriod     = 10 * time.Second
)
