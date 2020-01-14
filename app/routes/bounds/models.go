package bounds

//geographical boundary belonging to a larger bounded area
//swagger:model bounds
type boundsModel struct {
	//bounded it this bounds belongs to
	//required: true
	BoundId string `json:"bound_id"`
	//base64 encoded version of the bounds polyline
	//required: true
	Bounds string `json:"bounds"`
	//base64 encoded version of the bounds center point
	//required: true
	Centroid string `json:"centroid"`
}

//a list of bounds
//swagger:model boundsList
type boundsModelList struct {
	//items in the list
	//required: true
	Items []boundsModel `json:"items"`
}

//definition for a type of boundary
//swagger:model boundType
type boundTypeModel struct {
	//unique, identifiable name for the type
	//required: true
	TypeName string `json:"typeName"`
	//pretty, human readable display name for the type
	//required: true
	DisplayName string `json:"displayName"`
}

//A list of bounds types
//swagger:model boundTypeList
type boundTypeModelList struct {
	//items in the list
	//required: true
	Items []boundTypeModel `json:"items"`
}
