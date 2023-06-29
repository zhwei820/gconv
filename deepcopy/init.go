package deepcopy

import (
	"encoding/json"
	"reflect"

	"github.com/zhwei820/gconv/empty"
)

func CopyExported(dst interface{}, src interface{}) error { // only for exported member
	if res, err := json.Marshal(src); err != nil {
		return err
	} else {
		err = json.Unmarshal(res, dst)
		return err
	}
}

// only support 1 level
// user must assure src and dst are pointers, and have same field type
func SimpleCopyStruct(structDest, structSrc interface{}, ignoreEmpty ...bool) {
	// Use reflect to iterate fields and set values
	v := reflect.ValueOf(structSrc).Elem()
	typeOfP := v.Type()

	ignoreEmptyVal := len(ignoreEmpty) > 0 && (ignoreEmpty[0])

	m := map[string]interface{}{}
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if !ignoreEmptyVal || (ignoreEmptyVal && !empty.IsEmpty(field.Interface())) {
			m[typeOfP.Field(i).Name] = field.Interface()
		}
	}

	v = reflect.ValueOf(structDest).Elem()
	typeOfP = v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		name := typeOfP.Field(i).Name
		if val, ok := m[name]; ok {
			field.Set(reflect.ValueOf(val))
		}
	}
}
