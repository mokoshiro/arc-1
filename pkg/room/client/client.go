package client

import (
	"fmt"
	"log"
	"net/url"

	"net/http"

	"time"

	"encoding/binary"
	"encoding/json"

	"github.com/Bo0km4n/arc/pkg/room/api/message"
	"github.com/gorilla/websocket"
	"github.com/k0kubun/pp"
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

	c.WriteMessage(websocket.BinaryMessage, newPermissionRequest([]byte(`{
		"peers": [
			"bbbb",
			"cccc",
			"dddd"
		]}`)))
	c.WriteMessage(websocket.BinaryMessage, newRelayMessage([]string{
		"bbbb",
		"cccc",
		"dddd",
	}, []byte("hello, world")))
	defer c.Close()

	time.Sleep(10 * time.Second)
}

func newPermissionRequest(body []byte) []byte {
	b := []byte{0x00, 0x01}
	b = append(b, body...)
	pp.Println(b)
	return b
}

func newRelayMessage(dests []string, body []byte) []byte {
	req := []byte{0x00, 0x02}
	destReq := &message.Destinations{
		Peers: dests,
	}
	destBytes, _ := json.Marshal(destReq)
	destLen := uint32(len(destBytes))

	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, destLen)
	req = append(req, bytes...)
	req = append(req, destBytes...)
	req = append(req, body...)
	return req
}
