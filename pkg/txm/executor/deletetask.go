package executor

import "database/sql"

func deletePeerTx(tx *sql.Tx, req *DeletePeerRequest) error {
	exists, err := existPeer(tx, req.PeerID)
	if err != nil {
		return err
	}
	if !exists {
		return nil
	}
	if _, err := tx.Exec("DELETE FROM peer where peer_id = ?", req.PeerID); err != nil {
		return err
	}
	return tx.Commit()
}

func existPeer(tx *sql.Tx, peerID string) (bool, error) {
	var exists bool
	query := "SELECT exists (SELECT peer_id from peer where peer_id = ?)"
	err := tx.QueryRow(query, peerID).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return exists, nil
}
