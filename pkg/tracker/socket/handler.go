package socket

import (
	"fmt"

	"context"

	"time"

	"github.com/Bo0km4n/arc/pkg/tracker/h3"
	"github.com/Bo0km4n/arc/pkg/tracker/infra/db"
	"github.com/Bo0km4n/arc/pkg/tracker/msg"
)

func Greet(m *msg.GreetMessage) error {
	tx, err := db.MysqlPool.Begin()
	if err != nil {
		return err
	}
	statement := fmt.Sprintf("INSERT INTO peer(peer_id, addr, h3_hash, h3_resolution, location) VALUES (?, ?, ?, ?, ST_GeomFromText(?))")
	pointValue := fmt.Sprintf(`POINT(%f %f)`, m.Longitude, m.Latitude)
	ins, err := tx.Prepare(statement)
	if err != nil {
		tx.Rollback()
		return err
	}
	hash, err := h3.InvokeH3(14, m.Longitude, m.Latitude)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = ins.Exec(m.PeerID, m.Addr, hash, 14, pointValue)
	if err != nil {
		tx.Rollback()
		return err
	}

	conn := db.RedisPool.Get()
	if _, err := conn.Do("SETEX", m.PeerID, 86400, SOCK_ADDR); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func Tracking(m *msg.TrackingMessage) error {
	statement := fmt.Sprintf("UPDATE peer SET location = ST_GeomFromText(?) WHERE peer_id = ?")
	pointValue := fmt.Sprintf(`POINT(%f %f)`, m.Longitude, m.Latitude)
	stmt, err := db.MysqlPool.Prepare(statement)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(pointValue, m.PeerID)
	if err != nil {
		return err
	}
	return nil
}

func Lookup(m *msg.LookupRequest) (*msg.LookupResponse, error) {
	query := fmt.Sprintf(`
	SELECT peer_id, addr, ST_X(location), ST_Y(location) from peer WHERE ST_WITHIN(location, ST_BUFFER(POINT(?, ?), ?)) AND peer_id != ?
	`)
	radiusValue := convertRadius(m.Unit, m.Radius)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	rows, err := db.MysqlPool.QueryContext(ctx, query, m.Longitude, m.Latitude, radiusValue, m.PeerID)
	if err != nil {
		return nil, err
	}

	res := &msg.LookupResponse{
		Peers: make([]*msg.LookupResponseRow, 0),
	}
	for rows.Next() {
		r := &msg.LookupResponseRow{}
		if err := rows.Scan(&r.PeerID, &r.Addr, &r.Longitude, &r.Latitude); err != nil {
			return nil, err
		}
		res.Peers = append(res.Peers, r)
	}

	return res, nil
}

func convertRadius(unit string, radius float64) float64 {
	switch unit {
	case "m":
		radius /= 1000
	default:
		radius = radius
	}

	return (360 * radius) / 40075
}
