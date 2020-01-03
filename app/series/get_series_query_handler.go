package series

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/Masterminds/structable"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
	"time"

	"github.com/mdm373/ny-data-api/app/db"
	"github.com/mdm373/ny-data-api/app/router"
)

func mapRows(rows []structable.Recorder) []seriesRecord {
	values := make([]seriesRecord, len(rows))
	for i, item := range rows {
		values[i] = *item.Interface().(*seriesRecord)
	}
	return values
}

type granularity int32

type seriesRoute struct {
	config    seriesTypeRow
	recorders map[granularity]structable.Recorder
}

const (
	granularityDay granularity = iota
	granularityMonth
	granularityYear
)

var granularityParamMap = map[string]granularity{
	"day":   granularityDay,
	"month": granularityMonth,
	"year":  granularityYear,
}

func getTimeBound(r *http.Request, param string) (time.Time, error) {
	val, ok := r.URL.Query()[param]
	asTime := time.Date(100, 0, 0, 0, 0, 0, 0, time.UTC)
	var err error
	if ok {
		asTime, err = time.Parse(time.RFC3339, fmt.Sprintf("%sT00:00:00Z", val[0]))
		if err != nil {
			return time.Now(), errors.Wrap(err, fmt.Sprintf("failed to parse %s time", param))
		}
	}
	return asTime, nil
}
func getWhereFunc(r *http.Request) (structable.WhereFunc, error) {
	startTime, err := getTimeBound(r, "start")
	if err != nil {
		return nil, err
	}
	endTime, err := getTimeBound(r, "end")
	if err != nil {
		return nil, err
	}
	return func(desc structable.Describer, query squirrel.SelectBuilder) (squirrel.SelectBuilder, error) {
		return query.Where(squirrel.And{
			squirrel.GtOrEq{"timestamp": startTime.Format(time.RFC3339)},
			squirrel.LtOrEq{"timestamp": endTime.Format(time.RFC3339)},
		}), nil
	}, nil
}

func getSeriesQueryHandler(config []seriesTypeRow, connection db.Connection) router.RouteHandler {
	mapConfig := make(map[string]seriesRoute)
	for _, configItem := range config {
		mapConfig[configItem.TypeName] = seriesRoute{
			config: configItem,
			recorders: map[granularity]structable.Recorder{
				granularityDay:   connection.Bind(fmt.Sprintf("%s_day", configItem.TableName), seriesRecord{}),
				granularityMonth: connection.Bind(fmt.Sprintf("%s_month", configItem.TableName), seriesRecord{}),
				granularityYear:  connection.Bind(fmt.Sprintf("%s_year", configItem.TableName), seriesRecord{}),
			},
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		boundsType, ok := vars[idPathParam]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		granularityParam, ok := vars[granularityPathParam]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		granularityValue, ok := granularityParamMap[granularityParam]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		seriesTypeRoute, ok := mapConfig[boundsType]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		whereFunc, err := getWhereFunc(r)
		if err != nil {
			http.Error(w, fmt.Sprintf("query error: %+v", err), http.StatusInternalServerError)
			return
		}
		rows, err := structable.ListWhere(seriesTypeRoute.recorders[granularityValue], whereFunc)
		if err != nil {
			http.Error(w, fmt.Sprintf("db error: %+v", err), http.StatusInternalServerError)
			return
		}
		values := mapRows(rows)
		router.RespondWithJSON(w, values)
	}
}
