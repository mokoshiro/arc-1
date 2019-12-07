package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Bo0km4n/arc/pkg/broker/cmd/option"
	_ "github.com/go-sql-driver/mysql"
)

var (
	MysqlPool *sql.DB
)

func InitMysql() {
	host := option.Opt.MysqlHost
	port := option.Opt.MysqlPort
	user := option.Opt.MysqlUser
	password := option.Opt.MysqlPassword
	database := option.Opt.MysqlDatabase

	// url := "reckoner:-:HOGAHI@cloudsql(cdp-dev-206602:asia-east1:cdp-dev-taiwan-mysql-6ed0d135)/cdp?parseTime=true&charset=utf8mb4"
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4", user, password, host, port, database)

	c, err := sql.Open("mysql", url)
	if err != nil {
		log.Fatal(err)
	}
	if c == nil {
		log.Fatalf("Failed connect to mysql, url=%s", url)
	}
	c.SetMaxIdleConns(option.Opt.MysqlMaxIdleConns)
	c.SetMaxOpenConns(option.Opt.MysqlMaxOpenConns)

	MysqlPool = c
}
