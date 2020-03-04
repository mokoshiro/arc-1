package manager

import (
	"database/sql"
	"net/http"

	"github.com/garyburd/redigo/redis"
)

type ManagerServer struct {
	kvs        *redis.Pool
	db         *sql.DB
	httpClient *http.Client
}
