package series

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mdm373/ny-data-api/app/db"
	"github.com/mdm373/ny-data-api/app/router"
)

func AppendRoute(parent *mux.Router, connection db.Connection) error {
	boundConfig, err := getConfig(connection)
	if err != nil {
		return err
	}
	seriesRouter := parent.PathPrefix("/series").Subrouter()
	queryRoute := fmt.Sprintf("/query/{%s}/{%s}", idPathParam, granularityPathParam)
	router.AppendOptionedGetRoute(seriesRouter, queryRoute, getSeriesQueryHandler(boundConfig, connection))
	router.AppendOptionedGetRoute(seriesRouter, "/types/", getSeriesTypesHandler(boundConfig))
	return nil
}
