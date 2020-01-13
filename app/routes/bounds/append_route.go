package bounds

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mdm373/ny-data-api/app/db"
	"github.com/mdm373/ny-data-api/app/router"
)

// swagger:parameters listBoundsPaths
type paramsBounds struct {
	// name of the bounds type
	// in: path
	// required: true
	TypeName string `json:"type-name"`
}

var boundsParamsDef = paramsBounds{
	TypeName: "type-name",
}

func AppendRoute(parent *mux.Router, connection db.Connection) error {
	boundConfig, err := GetConfig(connection)
	if err != nil {
		return err
	}
	boundsRouter := parent.PathPrefix("/bounds").Subrouter()
	pathRoute := fmt.Sprintf("/paths/{%s}/", boundsParamsDef.TypeName)

	// swagger:route GET /bounds/paths/{type-name}/  bounds listBoundsPaths
	// get all paths for a given bounds type
	// Responses:
	//       200: boundsList
	//       500: error
	router.AppendOptionedGetRoute(boundsRouter, pathRoute, newGetBoundsHandler(boundConfig, connection))

	// swagger:route GET /bounds/types bounds listBoundsTypes
	// gets all available bounds data types
	// Responses:
	//       200: boundTypeList
	//       500: error
	router.AppendOptionedGetRoute(boundsRouter, "/types/", newGetBoundTypesHandler(boundConfig))
	return nil
}
