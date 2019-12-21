package bounds

import (
	"net/http"

	"github.com/mdm373/ny-data-api/app/router"
)

type BoundTypeRow struct {
	TypeName    string `json:"typeName"`
	DisplayName string `json:"displayName"`
}

func newGetBoundTypesHandler(config []BoundConfig) router.RouteHandler {
	var boundTypes []BoundTypeRow
	for _, configItem := range config {
		boundTypes = append(boundTypes, BoundTypeRow{
			TypeName:    configItem.Route,
			DisplayName: configItem.DisplayName,
		})
	}
	return func(w http.ResponseWriter, r *http.Request) {
		router.RespondWithJSON(w, boundTypes)
	}
}
