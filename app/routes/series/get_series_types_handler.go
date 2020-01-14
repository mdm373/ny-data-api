package series

import (
	"net/http"

	"github.com/mdm373/ny-data-api/app/router"
)

//Time series record set type definition
//swagger:model
type seriesType struct {
	//the series type name
	TypeName string `json:"typeName"`
	//human readable, pretty name of the type
	DisplayName string `json:"displayName"`
	// human readable, pretty name of the statistical value for this series
	ValueName string `json:"valueName"`
	// bound type that this series data relates to
	BoundType string `json:"boundType"`
	// least recent time on record for this time series
	Oldest string `json:"oldest"`
	// most recent time available for this time series
	Newest string `json:"newest"`
}

//List of series types
//swagger:model
type seriesTypeList struct {
	//items in this list
	Items []seriesType
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
		router.RespondWithJSON(w, seriesTypeList{Items: safeTypes})
	}
}
