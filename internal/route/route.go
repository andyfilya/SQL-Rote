package route 

import (
  "github.com/gorilla/mux"
)

func InitRouter(routes Routes) *mux.Router {
  router := mux.NewRouter()
  
  for _, route := range routes {
    router.Methods(route.Method).
      Path(route.Pattern).
      Name(route.Name).
      Handler(route.HandleFunc)
  }

  return router
}
