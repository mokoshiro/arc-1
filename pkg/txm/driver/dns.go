package driver

import "database/sql"

type executorDNSClient struct {
	db *sql.DB
}

func newExecutorDNSClient(db *sql.DB) *executorDNSClient {
	return &executorDNSClient{
		db: db,
	}
}
