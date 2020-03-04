package executor

import (
	"database/sql"

	"github.com/Bo0km4n/arc/pkg/txm/executor/schema"
)

func selectRowPeer(db *sql.DB, peerID string) (*schema.Peer, error) {
	query := `select peer_id, addr, credential, ST_X(location), ST_Y(location), created_at, updated_at from peer where peer_id = ?;`
	row := db.QueryRow(query, peerID)
	return deserializeRowToPeer(row)
}

func deserializeRowToPeer(row *sql.Row) (*schema.Peer, error) {
	peer := &schema.Peer{}
	err := row.Scan(
		&peer.PeerID,
		&peer.Addr,
		&peer.Credential,
		&peer.Longitude,
		&peer.Latitude,
		&peer.CreatedAt,
		&peer.UpdatedAt,
	)
	return peer, err
}
