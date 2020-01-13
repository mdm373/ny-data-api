package bounds

//geographical boundary belonging to a larger bounded area
//swagger:model bounds
type boundsModel struct {
	//bounded it this bounds belongs to
	BoundId string `json:"bound_id"`
	//base64 encoded version of the bounds polyline
	Bounds string `json:"bounds"`
	//base64 encoded version of the bounds center point
	Centroid string `json:"centroid"`
}

//a list of bounds
//swagger:model boundsList
type boundsModelList struct {
	//items in the list
	Items []boundsModel `json:"items"`
}

//definition for a type of boundary
//swagger:model boundType
type boundTypeModel struct {
	//unique, identifiable name for the type
	TypeName string `json:"typeName"`
	//pretty, human readable display name for the type
	DisplayName string `json:"displayName"`
}

//A list of bounds types
//swagger:model boundTypeList
type boundTypeModelList struct {
	//items in the list
	Items []boundTypeModel `json:"items"`
}
