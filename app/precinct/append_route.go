package precinct

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mdm373/ny-data-api/app/db"
	"net/http"
)

func AppendRoute(parent *mux.Router, connection db.Connection) {
	precinctByIDRouter := parent.PathPrefix(fmt.Sprintf("/precinct/{%s}", idPathParam)).Subrouter()
	precinctByIDRouter.HandleFunc("/bounds", newGetBoundsHandler(connection)).Methods(http.MethodGet, http.MethodOptions)
}
