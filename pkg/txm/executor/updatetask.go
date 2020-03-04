package executor

import (
	"context"
	"database/sql"
	"fmt"
)

type UpdateTask struct {
	tx    *sql.Tx
	state TaskState
}

func (ut *UpdateTask) prepare(ctx context.Context, req *UpdatePeerRequest) error {
	if ut.state == StateNone {
		// prepare statement
		query := `INSERT INTO peer (peer_id, addr, credential, location) VALUES(?, ?, ?, ST_GeomFromText(?)) ON DUPLICATE KEY UPDATE peer_id = ?`
		pointValue := fmt.Sprintf(`POINT(%f %f)`, req.Longitude, req.Latitude)

		stmt, err := ut.tx.Prepare(query)
		if err != nil {
			ut.tx.Rollback()
			return err
		}

		if _, err := stmt.Exec(req.PeerID, req.Addr, req.Credential, pointValue, req.PeerID); err != nil {
			ut.tx.Rollback()
			return err
		}
		// exec query
		ut.toReadyCommit()
	}
	return ut.tx.Commit()
}

func (ut *UpdateTask) commit() error {
	if ut.state == StateReadyCommit {
		return ut.tx.Commit()
	}
	return nil
}

func (ut *UpdateTask) rollback() error {
	if ut.canRollback() {
		return ut.tx.Rollback()
	}
	return nil
}

func (ut *UpdateTask) toReadyCommit() {
	ut.state = StateReadyCommit
}

func (ut *UpdateTask) toCommitted() {
	ut.state = StateCommitted
}

func (ut *UpdateTask) canRollback() bool {
	return ut.state == StateReadyCommit || ut.state == StateErrFailedInsert
}

func (ut *UpdateTask) toRollbacked() {
	ut.state = StateRollbacked
}
