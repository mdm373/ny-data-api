package series

import (
	"net/http"

	"github.com/mdm373/ny-data-api/app/router"
)

type seriesType struct {
	TypeName    string `json:"typeName"`
	DisplayName string `json:"displayName"`
	ValueName   string `json:"valueName"`
	BoundType   string `json:"boundType"`
	Oldest      string `json:"oldest"`
	Newest      string `json:"newest"`
}

func getSeriesTypesHandler(boundTypes []seriesTypeRow) router.RouteHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		var safeTypes []seriesType
		for _, item := range boundTypes {
			safeTypes = append(safeTypes, seriesType{
				TypeName:    item.TypeName,
				DisplayName: item.DisplayName,
				ValueName:   item.ValueName,
				BoundType:   item.BoundType,
				Oldest:      item.Oldest,
				Newest:      item.Newest,
			})
		}
		router.RespondWithJSON(w, safeTypes)
	}
}