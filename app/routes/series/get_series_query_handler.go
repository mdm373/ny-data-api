package series

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/Masterminds/structable"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
	"strings"
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

func getTimeStringFromRaw(val string) string {
	if strings.Index(val, "T") < 0 {
		return fmt.Sprintf("%sT00:00:00Z", val)
	}
	return val
}
func getTimeParamOrDefault(r *http.Request, param string, defaultTime string) (time.Time, error) {
	val, ok := r.URL.Query()[param]
	if ok {
		return time.Parse(time.RFC3339, getTimeStringFromRaw(val[0]))
	} else {
		return time.Parse(time.RFC3339, getTimeStringFromRaw(defaultTime))
	}
}
func getTimeBound(r *http.Request, param string, defaultTime string) (time.Time, error) {
	asTime, err := getTimeParamOrDefault(r, param, defaultTime)
	if err != nil {
		return time.Now(), errors.Wrap(err, fmt.Sprintf("failed to parse %s time", param))
	}
	return asTime, nil
}
func getWhereFunc(r *http.Request, aType seriesTypeRow) (structable.WhereFunc, error) {
	startTime, err := getTimeBound(r, pathParamsDef.Start, aType.Oldest)
	if err != nil {
		return nil, err
	}
	endTime, err := getTimeBound(r, pathParamsDef.End, aType.Newest)
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
		seriesType, ok := vars[pathParamsDef.SeriesType]
		if !ok {
			router.RespondWithError(w, "missing series type")
			return
		}
		granularityParam, ok := vars[pathParamsDef.Granularity]
		if !ok {
			router.RespondWithError(w, "missing granularity")
			return
		}
		granularityValue, ok := granularityParamMap[granularityParam]
		if !ok {
			router.RespondWithError(w, "invalid granularity")
			return
		}
		seriesTypeRoute, ok := mapConfig[seriesType]
		if !ok {
			router.RespondWithError(w, "invalid series type")
			return
		}
		whereFunc, err := getWhereFunc(r, seriesTypeRoute.config)
		if err != nil {
			router.RespondWithError(w, fmt.Sprintf("query error: %+v", err))
			return
		}
		rows, err := structable.ListWhere(seriesTypeRoute.recorders[granularityValue], whereFunc)
		if err != nil {
			router.RespondWithError(w, fmt.Sprintf("db error: %+v", err))
			return
		}
		values := mapRows(rows)
		router.RespondWithJSON(w, seriesRecordList{Items: values})
	}
}
