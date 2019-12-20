package bounds

import (
	"fmt"
	"github.com/Masterminds/structable"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/mdm373/ny-data-api/app/db"
	"github.com/mdm373/ny-data-api/app/router"
)

func mapRows(rows []structable.Recorder) []boundsRecord {
	values := make([]boundsRecord, len(rows))
	for i, item := range rows {
		values[i] = *item.Interface().(*boundsRecord)
	}
	return values
}

type boundsRoute struct {
	config   BoundConfig
	recorder structable.Recorder
}

func newGetBoundsHandler(config []BoundConfig, connection db.Connection) router.RouteHandler {
	mapConfig := make(map[string]boundsRoute)
	for _, configItem := range config {
		mapConfig[configItem.Route] = boundsRoute{
			config:   configItem,
			recorder: connection.Bind(configItem.TableName, boundsRecord{}),
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		boundsType, ok := vars[idPathParam]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		boundsTypeRoute, ok := mapConfig[boundsType]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		rows, err := structable.List(boundsTypeRoute.recorder, 1000, 0)
		if err != nil {
			http.Error(w, fmt.Sprintf("db error: %+v", err), http.StatusInternalServerError)
			return
		}
		values := mapRows(rows)
		router.RespondWithJSON(w, values)
	}
}
