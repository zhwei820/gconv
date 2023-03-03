package gconv

import (
	"reflect"
)

// Convert to type of toTypeValue
func ConvertTo(v interface{}, toTypeValue interface{}) interface{} {
	if toTypeValue == nil {
		return v
	}
	toTypeName := reflect.TypeOf(toTypeValue).String()
	return Convert(v, toTypeName)
}
