package precinct

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/mdm373/ny-data-api/app/db"
	"github.com/mdm373/ny-data-api/app/router"
	"net/http"
)

func newGetPrecinctsHandler(connection db.Connection) router.RouteHandler {
	query := squirrel.Select("precinct").Distinct().From(tableName)
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := connection.RunQuery(query)
		if err != nil {
			http.Error(w, fmt.Sprintf("db error: %+v", err), http.StatusInternalServerError)
		}
		slice := []string{}
		for rows.Next() {
			var precinctId string
			rows.Scan(&precinctId)
			slice = append(slice, precinctId)
		}
		router.RespondWithJSON(w, slice)

	}
}
