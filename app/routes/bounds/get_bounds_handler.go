package bounds

import (
	"fmt"
	"github.com/Masterminds/structable"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/mdm373/ny-data-api/app/db"
	"github.com/mdm373/ny-data-api/app/router"
)

type boundsRecord struct {
	Id       int    `stbl:"id"`
	BoundId  string `stbl:"bound_id"`
	Bounds   string `stbl:"bounds"`
	Centroid string `stbl:"centroid"`
}

func mapRecords(rows []structable.Recorder) []boundsModel {
	models := make([]boundsModel, len(rows))
	for i, item := range rows {
		record := *item.Interface().(*boundsRecord)
		models[i] = boundsModel{
			BoundId:  record.BoundId,
			Bounds:   record.Bounds,
			Centroid: record.Centroid,
		}
	}
	return models
}

type boundsRoute struct {
	config   BoundTypeRow
	recorder structable.Recorder
}

func newGetBoundsHandler(config []BoundTypeRow, connection db.Connection) router.RouteHandler {
	mapConfig := make(map[string]boundsRoute)
	for _, configItem := range config {
		mapConfig[configItem.TypeName] = boundsRoute{
			config:   configItem,
			recorder: connection.Bind(configItem.TableName, boundsRecord{}),
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		boundsType, ok := vars[boundsParamsDef.TypeName]
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
			router.RespondWithError(w, fmt.Sprintf("db error: %+v", err))
			return
		}
		models := mapRecords(rows)
		router.RespondWithJSON(w, boundsModelList{Items: models})
	}
}
