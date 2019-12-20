package bounds

const (
	idPathParam = "id"
)

type boundsRecord struct {
	Id       int    `stbl:"id" json:"id"`
	Precinct string `stbl:"bound_id" json:"bound_id"`
	Bounds   string `stbl:"bounds" json:"bounds"`
	Centroid string `stbl:"centroid" json:"centroid"`
}
