package executor

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func mysql(host, port, user, password, database string, maxIdleConns, maxOpenConns int) *sql.DB {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4", user, password, host, port, database)

	c, err := sql.Open("mysql", url)
	if err != nil {
		log.Fatal(err)
	}
	if c == nil {
		log.Fatalf("Failed connect to mysql, url=%s", url)
	}
	c.SetMaxIdleConns(maxIdleConns)
	c.SetMaxOpenConns(maxOpenConns)
	return c
}
