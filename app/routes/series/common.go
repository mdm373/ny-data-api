package series

//geographical boundary belonging to a larger bounded area
//swagger:model seriesRecord
type seriesRecord struct {
	//sequential unique identifier of this data point in series
	//required: true
	Id int `stbl:"id" json:"id"`
	// id of the boundry instance this record belongs to
	//required: true
	BoundId string `stbl:"bound_id" json:"bound_id"`
	// timestamp for the beginning of this series record's duration
	//required: true
	Timestamp string `stbl:"timestamp" json:"timestamp"`
	// statistical value for this series record
	//required: true
	Value string `stbl:"value" json:"value"`
}

//A list of series records
//swagger:model seriesRecordList
type seriesRecordList struct {
	// the items in this list
	//required: true
	Items []seriesRecord `json:"items"`
}
