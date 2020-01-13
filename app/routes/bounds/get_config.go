package bounds

import (
	"github.com/Masterminds/structable"
	"github.com/mdm373/ny-data-api/app/db"
)

func mapBoundTypeRows(rows []structable.Recorder) []BoundTypeRow {
	values := make([]BoundTypeRow, len(rows))
	for i, item := range rows {
		values[i] = *item.Interface().(*BoundTypeRow)
	}
	return values
}

type BoundTypeRow struct {
	TypeName    string `stbl:"type_name"`
	DisplayName string `stbl:"display_name"`
	TableName   string `stbl:"table_name"`
}

var GetConfig = func(connection db.Connection) ([]BoundTypeRow, error) {
	recorder := connection.Bind("bounds_types", BoundTypeRow{})
	boundTypes, err := structable.List(recorder, 1000, 0)
	if err != nil {
		return nil, err
	}
	return mapBoundTypeRows(boundTypes), nil
}
