package executor

import (
	"context"
	"database/sql"
	"fmt"
)

type PutTask struct {
	tx    *sql.DB
	state TaskState
}

type TaskState string

const (
	StateNone            = "None"
	StateReadyCommit     = "ReadyCommit"
	StateCommitted       = "Committed"
	StateRollbacked      = "Rollbacked"
	StateErrFailedInsert = "ErrFailedInsert"
	StateErrFailedFetch  = "ErrFailedFetch"
)

func (pt *PutTask) prepare(ctx context.Context, req *PreparePutPeerRequest) error {
	if pt.state == StateNone {
		exists, err := pt.isExistPeer(req.PeerID)
		if err != nil {
			return err
		}
		if exists {
			return nil
		}
		statement := fmt.Sprintf("INSERT INTO peer(peer_id, addr, credential, location) VALUES (?, ?, ?, ST_GeomFromText(?))")
		pointValue := fmt.Sprintf(`POINT(%f %f)`, req.Longitude, req.Latitude)
		ins, err := pt.tx.Prepare(statement)
		if err != nil {
			return err
		}
		_, err = ins.Exec(req.PeerID, req.Addr, req.Credential, pointValue)
		if err != nil {
			return err
		}
	}
	return nil
}

func (pt *PutTask) isExistPeer(peerID string) (bool, error) {
	var exists bool
	query := "SELECT exists (SELECT peer_id from peer where peer_id = ?)"
	err := pt.tx.QueryRow(query, peerID).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return exists, nil
}
