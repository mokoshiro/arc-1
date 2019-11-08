package socket

import (
	"fmt"

	"github.com/Bo0km4n/arc/pkg/tracker/h3"
	"github.com/Bo0km4n/arc/pkg/tracker/infra/db"
)

func greet(m *greetMessage) error {
	statement := fmt.Sprintf("INSERT INTO peer(peer_id, h3_hash, h3_resolution, latitude, longitude) VALUES (?, ?, ?, ?, ?)")

	ins, err := db.MysqlPool.Prepare(statement)
	if err != nil {
		return err
	}
	hash, err := h3.InvokeH3(14, m.longitude, m.latitude)
	if err != nil {
		return err
	}
	_, err = ins.Exec(m.peerID, hash, 14, m.latitude, m.longitude)
	if err != nil {
		return err
	}

	conn := db.RedisPool.Get()
	if _, err := conn.Do("SETEX", m.peerID, SOCK_ADDR); err != nil {
		return err
	}
	return nil
}
