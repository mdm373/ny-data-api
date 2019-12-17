package precinct

const (
	tableName   = "nypd_precinct_bounds"
	idPathParam = "id"
)

var emptyPrecinctBounds = precinctBounds{}

type precinctBounds struct {
	Id       int    `stbl:"id" json:"id"`
	Precinct string `stbl:"precinct" json:"precinct"`
	Bounds   string `stbl:"bounds" json:"bounds"`
}

type precinct struct {
	Precinct string `stbl:precinct json: precinct`
}

var emptyPrecinct = precinct{}
