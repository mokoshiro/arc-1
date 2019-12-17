package client

import (
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"net/http"

	"encoding/binary"
	"encoding/json"

	"github.com/Bo0km4n/arc/pkg/room/api/message"
	"github.com/gorilla/websocket"
	"github.com/k0kubun/pp"
)

type input struct {
	ID         string   `json:"id"`
	Credential string   `json:"credential"`
	Host       string   `json:"host"`
	Permission []string `json:"permission"`
}

func Run(jsonin string) {

	pp.Println(jsonin)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	in := &input{}
	json.Unmarshal([]byte(jsonin), in)

	pp.Println(in)
	header := http.Header{}
	header.Add("X-ARC-PEER-ID", in.ID)
	header.Add("X-ARC-PEER-CREDENTIAL", in.Credential)
	u := url.URL{Scheme: "ws", Host: in.Host, Path: "/bind"}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), header)
	if err != nil {
		log.Fatal(err)
	}

	c.WriteMessage(websocket.BinaryMessage, newPermissionRequest(in))
	c.WriteMessage(websocket.BinaryMessage, newRelayMessage([]string{
		"bbbb",
	}, []byte("hello, world")))
	defer c.Close()

	go func() {
		for {
			mt, m, err := c.ReadMessage()
			if err != nil {
				log.Fatal(err)
			}
			if mt != websocket.TextMessage {
				log.Printf("recv: %s", m)
			}
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-interrupt:
			log.Println("interrupt")
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(3 * time.Second)
			return
		case <-ticker.C:
			c.WriteMessage(websocket.TextMessage, []byte("hello"))
		}
	}
}

func newPermissionRequest(in *input) []byte {
	b := []byte{0x00, 0x01}
	jsonb, _ := json.Marshal(&message.PermissionRequest{
		Peers: in.Permission,
	})
	b = append(b, jsonb...)
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
