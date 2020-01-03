package bounds

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mdm373/ny-data-api/app/db"
	"github.com/mdm373/ny-data-api/app/router"
)

func AppendRoute(parent *mux.Router, connection db.Connection) error {
	boundConfig, err := GetConfig(connection)
	if err != nil {
		return err
	}
	boundsRouter := parent.PathPrefix("/bounds").Subrouter()
	pathRoute := fmt.Sprintf("/paths/{%s}/", idPathParam)
	router.AppendOptionedGetRoute(boundsRouter, pathRoute, newGetBoundsHandler(boundConfig, connection))
	router.AppendOptionedGetRoute(boundsRouter, "/types/", newGetBoundTypesHandler(boundConfig))
	return nil
}
