package bounds

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mdm373/ny-data-api/app/db"
	"github.com/mdm373/ny-data-api/app/router"
)

func AppendRoute(parent *mux.Router, connection db.Connection) error {
	boundConfig, err := GetConfig()
	if err != nil {
		return err
	}
	boundsRouter := parent.PathPrefix(fmt.Sprintf("/bounds/{%s}", idPathParam)).Subrouter()
	router.AppendOptionedGetRoute(boundsRouter, "/", newGetBoundsHandler(boundConfig, connection))
	return nil
}
