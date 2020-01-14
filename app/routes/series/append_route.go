package series

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mdm373/ny-data-api/app/db"
	"github.com/mdm373/ny-data-api/app/router"
)

// swagger:parameters querySeries
type pathParams struct {
	// granularity level of the series data (year | month | day)
	// in: path
	// required: true
	Granularity string `json:"granularity"`
	// the type name of the time series data
	// in: path
	// required: true
	SeriesType string `json:"type-name"`
	// start date bounds (YYY-MM-DD). Defaults to series type's oldest date
	// in: query
	// required: false
	Start string `json:"start"`
	// start date bounds (YYYY-MM-DD). Defaults to series type's newest date
	// in: query
	// required: false
	End string `json:"end"`
}

var pathParamsDef = pathParams{
	SeriesType:  "seriesType",
	Granularity: "granularity",
	Start:       "start",
	End:         "end",
}

func AppendRoute(parent *mux.Router, connection db.Connection) error {
	boundConfig, err := getConfig(connection)
	if err != nil {
		return err
	}
	seriesRouter := parent.PathPrefix("/series").Subrouter()
	queryRoute := fmt.Sprintf("/query/{%s}/{%s}", pathParamsDef.SeriesType, pathParamsDef.Granularity)

	// swagger:route GET /series/query/{type-name}/{granularity}/ series querySeries
	// query time series records for a given type and granularity level
	// Responses:
	//       200: seriesRecordList
	//       500: error
	router.AppendOptionedGetRoute(seriesRouter, queryRoute, getSeriesQueryHandler(boundConfig, connection))
	// swagger:route GET /series/types/ series listSeriesTypes
	// list the available time series record types and their associated information
	// Responses:
	//       200: seriesTypeList
	//       500: error
	router.AppendOptionedGetRoute(seriesRouter, "/types/", getSeriesTypesHandler(boundConfig))
	return nil
}
