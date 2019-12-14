package db

import (
	"github.com/Masterminds/squirrel"
	"reflect"
)

func FieldEq(record interface{}, fieldName string, fieldValue interface{}) squirrel.Eq {
	fields := reflect.TypeOf(record)
	found := ""
	if field, ok := fields.FieldByName(fieldName); ok {
		found = field.Tag.Get("stbl")
	}
	return squirrel.Eq{found: fieldValue}
}
