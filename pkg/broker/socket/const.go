package socket

import (
	"fmt"
	"os"
	"time"

	"github.com/Bo0km4n/arc/pkg/broker/cmd/option"
	"github.com/gorilla/websocket"
)

var (
	upgrader  = websocket.Upgrader{}
	conns     = map[string]*Sock{}
	SOCK_ADDR = fmt.Sprintf("%s:%d", os.Getenv("BROKER_IP"), option.Opt.Port)
)

const (
	writeWait      = 10 * time.Second
	readLimit      = 60 * time.Second
	maxMessageSize = 512
	pongWait       = 60 * time.Second
	pingPeriod     = 10 * time.Second
)
