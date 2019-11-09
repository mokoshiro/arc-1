package socket

import (
	"fmt"

	"github.com/Bo0km4n/arc/pkg/tracker/h3"
	"github.com/Bo0km4n/arc/pkg/tracker/infra/db"
)

func greet(m *greetMessage) error {
	statement := fmt.Sprintf("INSERT INTO peer(peer_id, h3_hash, h3_resolution, point) VALUES (?, ?, ?, ?)")
	pointValue := fmt.Sprintf(`ST_GeomFromText('POINT(%f %f)')`, m.longitude, m.latitude)
	ins, err := db.MysqlPool.Prepare(statement)
	if err != nil {
		return err
	}
	hash, err := h3.InvokeH3(14, m.longitude, m.latitude)
	if err != nil {
		return err
	}
	_, err = ins.Exec(m.peerID, hash, 14, pointValue)
	if err != nil {
		return err
	}

	conn := db.RedisPool.Get()
	if _, err := conn.Do("SETEX", m.peerID, SOCK_ADDR); err != nil {
		return err
	}
	return nil
}

func tracking(m *trackingMessage) error {
	statement := fmt.Sprintf("UPDATE peer SET point = ? WHERE peer_id = ?")
	pointValue := fmt.Sprintf(`ST_GeomFromText('POINT(%f %f)')`, m.longitude, m.latitude)
	stmt, err := db.MysqlPool.Prepare(statement)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(pointValue, m.peerID)
	if err != nil {
		return err
	}
	return nil
}
