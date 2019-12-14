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

const idPathParam = "id"

type precinct struct {
	Sector     string `stbl:"sector" json:"sector"`
	PrecinctId string `stbl:"pct" json:"precinctId"`
	Geom       string `stbl:"the_geom" json:"geom"`
	Phase      string `stbl:"phase" json:"phase"`
}

func wherePrecinctId(precinctId string) structable.WhereFunc {
	return func(d structable.Describer, q squirrel.SelectBuilder) (squirrel.SelectBuilder, error) {
		return q.Where(db.FieldEq(precinct{}, "PrecinctId", precinctId)).Limit(10), nil
	}
}

func mapRows(rows []structable.Recorder) []precinct {
	values := make([]precinct, len(rows))
	for i, item := range rows {
		values[i] = *item.Interface().(*precinct)
	}
	return values
}

func newGetBoundsHandler(connection db.Connection) router.RouteHandler {
	recorder := connection.Bind(tableName, new(precinct))
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
