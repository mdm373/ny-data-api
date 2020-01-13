package bounds

import (
	"net/http"

	"github.com/mdm373/ny-data-api/app/router"
)

func newGetBoundTypesHandler(boundTypes []BoundTypeRow) router.RouteHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		var safeTypes []boundTypeModel
		for _, item := range boundTypes {
			safeTypes = append(safeTypes, boundTypeModel{
				TypeName:    item.TypeName,
				DisplayName: item.DisplayName,
			})
		}
		router.RespondWithJSON(w, boundTypeModelList{Items: safeTypes})
	}
}
