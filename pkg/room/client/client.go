package client

import (
	"fmt"
	"log"
	"net/url"

	"net/http"

	"github.com/gorilla/websocket"
)

func Run(peerID, credential, host string) {
	fmt.Println(peerID, credential, host)
	header := http.Header{}
	header.Add("X-ARC-PEER-ID", peerID)
	header.Add("X-ARC-PEER-CREDENTIAL", credential)
	u := url.URL{Scheme: "ws", Host: host, Path: "/bind"}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), header)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
}
