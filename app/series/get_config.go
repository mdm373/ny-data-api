package series

import (
	"github.com/Masterminds/structable"
	"github.com/mdm373/ny-data-api/app/db"
)

func mapSeriesTypeRows(rows []structable.Recorder) []seriesTypeRow {
	values := make([]seriesTypeRow, len(rows))
	for i, item := range rows {
		values[i] = *item.Interface().(*seriesTypeRow)
	}
	return values
}

type seriesTypeRow struct {
	TypeName    string `stbl:"type_name"`
	DisplayName string `stbl:"display_name"`
	TableName   string `stbl:"table_name"`
	BoundType   string `stbl:"bound_type"`
	ValueName   string `stbl:"value_name"`
	Oldest      string `stbl:"oldest"`
	Newest      string `stbl:"newest"`
}

var getConfig = func(connection db.Connection) ([]seriesTypeRow, error) {
	recorder := connection.Bind("series_types", seriesTypeRow{})
	boundTypes, err := structable.List(recorder, 1000, 0)
	if err != nil {
		return nil, err
	}
	return mapSeriesTypeRows(boundTypes), nil
}
