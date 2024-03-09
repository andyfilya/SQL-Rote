package route

import (
	"net/http"

	"github.com/sqlroute/internal/handlers"
)

type Route struct {
  Name string
  Method string
  Pattern string
  HandleFunc http.HandlerFunc
}

type Routes []Route

func InitRoutes(newhandlers *handlers.Handler) Routes {
  var routes = Routes{
    Route{
      "GetTables",
      "GET",
      "/",
      newhandlers.AllTables,
    },
    Route{
      "GetTableId",
      "GET",
      "/table/{tableId}",
      newhandlers.TableInformation,
    },
    Route{
      "RowInformation",
      "GET",
      "/table?limit={limit}&offset={offset}",
      newhandlers.LimitTables,
    },
    Route{
      "NewTable",
      "PUT",
      "/table",
      newhandlers.NewTable,
    },
    Route{
      "UpdateTable",
      "POST",
      "/table/{tableId}",
      newhandlers.UpdateTable,
    },
    Route{
      "DeleteRoute",
      "DELETE",
      "/table/{tableId}",
      newhandlers.DeleteRow,
    },
  }

  return routes
}
