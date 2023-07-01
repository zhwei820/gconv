// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gconv

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

// apiVal is used for type assert api for Val().
type apiVal interface {
	Val() interface{}
}

// Dump prints variables <i...> to stdout with more manually readable.
func Dump(i ...interface{}) {
	s := Export(i...)
	if s != "" {
		fmt.Println(s)
	}
}

const EXPORT_KEY = "GOEXPORT"

func SetExportExpand(expand bool) {
	if expand {
		os.Setenv(EXPORT_KEY, "1")
	} else {
		os.Setenv(EXPORT_KEY, "")
	}
}

// Export returns variables <i...> as a string with more manually readable.
func Export(i ...interface{}) string {
	if len(i) == 0 {
		return ""
	}
	goexport := os.Getenv(EXPORT_KEY)
	if goexport != "" {
		return export(i...)
	} else {
		return String(i[0])
	}
}

func export(i ...interface{}) string {
	buffer := bytes.NewBuffer(nil)
	for _, value := range i {
		switch r := value.(type) {
		case []byte:
			buffer.Write(r)
		case string:
			buffer.WriteString(r)
		default:
			var (
				reflectValue = reflect.ValueOf(value)
				reflectKind  = reflectValue.Kind()
			)
			for reflectKind == reflect.Ptr {
				reflectValue = reflectValue.Elem()
				reflectKind = reflectValue.Kind()
			}
			switch reflectKind {
			case reflect.Slice, reflect.Array:
				value = Interfaces(value)
			case reflect.Map:
				value = Map(value)
			case reflect.Struct:
				// converted := false
				// if r, ok := value.(apiVal); ok {
				// 	if result := r.Val(); result != nil {
				// 		value = result
				// 		converted = true
				// 	}
				// }
				// if !converted {
				// 	if r, ok := value.(apiMapStrAny); ok {
				// 		if result := r.MapStrAny(); result != nil {
				// 			value = result
				// 			converted = true
				// 		}
				// 	}
				// }
				// if !converted {
				// 	if r, ok := value.(apiString); ok {
				// 		value = r.String()
				// 	}
				// }
			}
			encoder := json.NewEncoder(buffer)
			encoder.SetEscapeHTML(false)
			encoder.SetIndent("", "\t")
			if err := encoder.Encode(value); err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			}
		}
	}
	return buffer.String()
}
