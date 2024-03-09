package main

import (
	"github.com/sirupsen/logrus"
	"github.com/sqlroute/internal/db"
	"github.com/sqlroute/internal/handlers"
	"github.com/sqlroute/internal/server"
  "github.com/sqlroute/internal/route"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.InitDB()
	if err != nil {
		logrus.Errorf("error while init db in main function %v", err)
		return
	}
	err = db.Ping()
	if err != nil {
		logrus.Errorf("error while first time ping new mysql database %v", err)
		return
	}
  handlers := &handlers.Handler{
    DB: db,
  }
  routes := route.InitRoutes(handlers)
	err = server.StartWebServer("8081", routes)
	if err != nil {
		logrus.Errorf("server ended with error in main function %v", err)
		return
  }
}
