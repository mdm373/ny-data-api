package precinct

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mdm373/ny-data-api/app/db"
	"github.com/mdm373/ny-data-api/app/router"
)

func AppendRoute(parent *mux.Router, connection db.Connection) {
	precinctByIDRouter := parent.PathPrefix(fmt.Sprintf("/precinct/{%s}", idPathParam)).Subrouter()
	router.AppendOptionedGetRoute(precinctByIDRouter, "/bounds/", newGetPrecinctBoundsHandler(connection))
	router.AppendOptionedGetRoute(parent, "/precinct/", newGetPrecinctsHandler(connection))
	router.AppendOptionedGetRoute(parent, "/precinct/bounds/", newGetAllPrecinctBoundsHandler(connection))
}
