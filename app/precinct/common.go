package precinct

const (
	tableName   = "NYPD_SECTORS"
	idPathParam = "id"
)

var emptyPrecinct = precinct{}

type precinct struct {
	Sector     string `stbl:"sector" json:"sector"`
	PrecinctId string `stbl:"pct" json:"precinctId"`
	Geom       string `stbl:"the_geom" json:"geom"`
	Phase      string `stbl:"phase" json:"phase"`
}
