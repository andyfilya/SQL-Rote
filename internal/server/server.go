package server

import (
	"context"
	"net/http"
  
  "github.com/sirupsen/logrus"
	"github.com/sqlroute/internal/route"
)

func StartWebServer(port string, routes route.Routes) error {
  r := route.InitRouter(routes)

  server := http.Server{
    Addr: "localhost:"+port,
    Handler: r,
  }

  defer func() {
    server.Shutdown(context.Background())
  }()
  
  if err := server.ListenAndServe(); err != nil {
    logrus.Infof("server ended with error")
    return err
  }
  return nil
}
