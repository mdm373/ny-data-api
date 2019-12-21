package bounds

import (
	"net/http"

	"github.com/mdm373/ny-data-api/app/router"
)

type safeBoundType struct {
	TypeName    string `json:"typeName"`
	DisplayName string `json:"displayName"`
}

func newGetBoundTypesHandler(boundTypes []BoundTypeRow) router.RouteHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		var safeTypes []safeBoundType
		for _, item := range boundTypes {
			safeTypes = append(safeTypes, safeBoundType{
				TypeName:    item.TypeName,
				DisplayName: item.DisplayName,
			})
		}
		router.RespondWithJSON(w, safeTypes)
	}
}
