package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

func DSN() string {
	dsn := "root:love@tcp(localhost:3306)/realdb?"
	dsn += "&charset=utf8"
	dsn += "&interpolateParams=true"

	return dsn
}

func InitDB() (*sql.DB, error) {
	dsn := DSN()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logrus.Infof("error while open mysql db in init db function %v", err)
		return nil, err
	}
	return db, nil
}
