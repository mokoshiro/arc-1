package client

import (
	"fmt"
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
	Mode       string   `json:"mode"`
	Credential string   `json:"credential"`
	Host       string   `json:"host"`
	Permission []string `json:"permission"`
	Span       int      `json:"span"`
	Frequency  int      `json:"frequency"`
	Chunk      int      `json:"chunk"`
}

func Run(jsonin string) {
	in := &input{}
	json.Unmarshal([]byte(jsonin), in)

	pp.Println(in)
	switch in.Mode {
	case "sender":
		sender(in)
	case "receiver":
		receiver(in)
	case "bench_sender":
		benchSender(in)
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

func sender(in *input) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	header := http.Header{}
	header.Add("X-ARC-PEER-ID", in.ID)
	header.Add("X-ARC-PEER-CREDENTIAL", in.Credential)
	u := url.URL{Scheme: "ws", Host: in.Host, Path: "/bind"}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), header)
	if err != nil {
		log.Fatal(err)
	}

	c.WriteMessage(websocket.BinaryMessage, newPermissionRequest(in))
	defer c.Close()

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
			c.WriteMessage(websocket.BinaryMessage, newRelayMessage(in.Permission, make([]byte, 512)))
			c.WriteMessage(websocket.TextMessage, []byte(""))
		}
	}
}

func benchSender(in *input) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	header := http.Header{}
	header.Add("X-ARC-PEER-ID", in.ID)
	header.Add("X-ARC-PEER-CREDENTIAL", in.Credential)
	u := url.URL{Scheme: "ws", Host: in.Host, Path: "/bind"}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), header)
	if err != nil {
		log.Fatal(err)
	}

	c.WriteMessage(websocket.BinaryMessage, newPermissionRequest(in))
	defer c.Close()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	benchTicker := time.NewTicker(time.Minute * time.Duration(in.Span))
	writeBytes := 0
	defer func() {
		log.Printf("writed bytes=%d", writeBytes)
	}()
	defer benchTicker.Stop()
	go func() {
		for {
			if err := c.WriteMessage(websocket.BinaryMessage, newRelayMessage(in.Permission, make([]byte, in.Chunk))); err != nil {
				log.Fatal(err)
			}
			writeBytes += in.Chunk
			time.Sleep(time.Duration(in.Frequency) * time.Millisecond)
			c.SetWriteDeadline(time.Now().Add(time.Second * 3))
			fmt.Fprintf(os.Stdout, "\rwrite bytes: %d", writeBytes)
		}
	}()

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
		// case <-ticker.C:
		// 	c.WriteMessage(websocket.TextMessage, []byte(""))
		case <-benchTicker.C:
			log.Printf("pasted %d minutes", in.Span)
			time.Sleep(3 * time.Second)
			return
		}
	}
}

func receiver(in *input) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	header := http.Header{}
	header.Add("X-ARC-PEER-ID", in.ID)
	header.Add("X-ARC-PEER-CREDENTIAL", in.Credential)
	u := url.URL{Scheme: "ws", Host: in.Host, Path: "/bind"}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), header)
	if err != nil {
		log.Fatal(err)
	}

	readBytes := 0
	go func() {
		for {
			mt, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			if mt == 2 {
				readBytes += len(message)
				fmt.Fprintf(os.Stdout, "\rreceive bytes: %d", readBytes)
			}
		}
	}()

	c.WriteMessage(websocket.BinaryMessage, newPermissionRequest(in))
	defer c.Close()
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
			c.WriteMessage(websocket.TextMessage, []byte(""))
		}
	}
}
