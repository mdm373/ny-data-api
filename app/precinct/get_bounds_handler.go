package precinct

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/Masterminds/structable"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/mdm373/ny-data-api/app/db"
	"github.com/mdm373/ny-data-api/app/router"
)

func wherePrecinctId(precinctId string) structable.WhereFunc {
	return func(d structable.Describer, q squirrel.SelectBuilder) (squirrel.SelectBuilder, error) {
		return q.Where(db.FieldEq(emptyPrecinctBounds, "Precinct", precinctId)).Limit(100), nil
	}
}

func mapRows(rows []structable.Recorder) []precinctBounds {
	values := make([]precinctBounds, len(rows))
	for i, item := range rows {
		values[i] = *item.Interface().(*precinctBounds)
	}
	return values
}

func newGetAllPrecinctBoundsHandler(connection db.Connection) router.RouteHandler {
	recorder := connection.Bind(tableName, emptyPrecinctBounds)
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := structable.List(recorder, 1000, 0)
		if err != nil {
			http.Error(w, fmt.Sprintf("db error: %+v", err), http.StatusInternalServerError)
			return
		}
		values := mapRows(rows)
		router.RespondWithJSON(w, values)
	}
}

func newGetPrecinctBoundsHandler(connection db.Connection) router.RouteHandler {
	recorder := connection.Bind(tableName, emptyPrecinctBounds)
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if precinctId, ok := vars[idPathParam]; ok {
			rows, err := structable.ListWhere(recorder, wherePrecinctId(precinctId))
			if err != nil {
				http.Error(w, fmt.Sprintf("db error: %+v", err), http.StatusInternalServerError)
				return
			}
			values := mapRows(rows)
			router.RespondWithJSON(w, values)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
	}
}
