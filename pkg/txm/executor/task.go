package executor

import "database/sql"

type PutTask struct {
	tx    *sql.Tx
	state TaskState
}

type TaskState string

const (
	ReadyCommit     = "ReadyCommit"
	Commited        = "Commited"
	Rollbacked      = "Rollbacked"
	ErrFailedInsert = "ErrFailedInsert"
	ErrFailedFetch  = "ErrFailedFetch"
)
