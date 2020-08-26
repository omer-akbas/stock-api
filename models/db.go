package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/benthor/clustersql"
	"github.com/go-sql-driver/mysql"
	"github.com/omer-akbas/stock-api/config"
)

var db *sql.DB

func Init() {
	var err error
	db, err = connString()
	if err != nil {
		log.Println("db connection errors. ", err.Error())
	}
}

func connString() (*sql.DB, error) {
	var err error
	cfg := config.Start(os.Args[1])
	var q = cfg.Mysql
	hostCount := len(q.Host)

	if hostCount == 0 {
		log.Fatal("fatal err: DB host name is null")
		return db, fmt.Errorf("DB host name is null")
	} else if hostCount == 1 {
		log.Println("conn: ONE connection")
		connQuery := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", q.User, q.Password, q.Host[0], q.Port, q.Database)
		db, err = sql.Open("mysql", connQuery)
	} else if hostCount > 1 {
		log.Println("conn: CLUSTER connection")
		mysqlDriver := mysql.MySQLDriver{}
		// err := mysql.SetLogger(mylogger)
		clusterDriver := clustersql.NewDriver(mysqlDriver)

		for i := 1; i <= hostCount; i++ {
			cQ := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", q.User, q.Password, q.Host[i-1], q.Port, q.Database)
			cN := fmt.Sprintf("galera%d", i)
			clusterDriver.AddNode(cN, cQ)
		}

		sql.Register("myCluster", clusterDriver)

		db, err = sql.Open("myCluster", "whatever")
	}
	return db, err
}
