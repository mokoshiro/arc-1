package middleware

import (
	"fmt"

	"github.com/Bo0km4n/arc/pkg/room/infra/db"
)

func ValidateCredential(peerID, credential string) error {
	query := fmt.Sprintf("SELECT id from peer WHERE peer_id=? and credential=?")
	row := db.MysqlPool.QueryRow(query, peerID, credential)
	return row.Scan()
}
