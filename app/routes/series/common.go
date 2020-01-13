package series

const (
	idPathParam          = "id"
	granularityPathParam = "granularity"
)

type seriesRecord struct {
	Id        int    `stbl:"id" json:"id"`
	BoundId   string `stbl:"bound_id" json:"bound_id"`
	Timestamp string `stbl:"timestamp" json:"timestamp"`
	Value     string `stbl:"value" json:"value"`
}
