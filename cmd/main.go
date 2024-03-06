package main

import (
	"github.com/sirupsen/logrus"
	"github.com/sqlroute/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mysqldb, err := db.InitDB()
	if err != nil {
		logrus.Errorf("error while init db in main function %v", err)
		return
	}
	err = mysqldb.Db.Ping()
	if err != nil {
		logrus.Errorf("error while first time ping new mysql database %v", err)
		return
	}
}
