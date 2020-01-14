package series

import (
	"net/http"

	"github.com/mdm373/ny-data-api/app/router"
)

//Time series record set type definition
//swagger:model
type seriesType struct {
	//the series type name
	//required: true
	TypeName string `json:"typeName"`
	//human readable, pretty name of the type
	//required: true
	DisplayName string `json:"displayName"`
	// human readable, pretty name of the statistical value for this series
	//required: true
	ValueName string `json:"valueName"`
	// bound type that this series data relates to
	//required: true
	BoundType string `json:"boundType"`
	// least recent time on record for this time series
	//required: true
	Oldest string `json:"oldest"`
	// most recent time available for this time series
	//required: true
	Newest string `json:"newest"`
}

//List of series types
//swagger:model
type seriesTypeList struct {
	//items in this list
	//required: true
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
