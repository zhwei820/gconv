package gin_bind

import (
	"context"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/zhwei820/errors"
)

func BindQuery(ctx context.Context, g *gin.Context, req interface{}) error {
	value := reflect.ValueOf(req)
	if value.Type().Kind() != reflect.Ptr {
		return errors.Errorf("req must be a pointer")
	}

	if err := g.BindQuery(req); err != nil {
		return err
	}
	doResetNil(value, g) // set form field to null when field is empty
	return nil
}

func doResetNil(value reflect.Value, g *gin.Context) {
	if value.IsNil() {
		return
	}

	v := value.Elem() // acquire value referenced by pointer
	if v.Type().Kind() != reflect.Struct {
		return
	}
	typeOfS := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fieldValV := v.Field(i)
		if fieldValV.Type().Kind() == reflect.Ptr {
			// fmt.Println("fieldVal.Type.Name()", reflect.TypeOf(fieldValV.Interface()).Elem())
			// fmt.Println("fieldValV.IsNil", fieldValV.IsNil())

			if fieldValV.IsNil() {
				continue
			}

			fieldVal := typeOfS.Field(i)
			formTag := getStructTag(fieldVal, "form")
			if formTag != "" {
				if g.Query(formTag) == "" {
					fieldValV.Set(reflect.NewAt(reflect.TypeOf(fieldValV.Interface()).Elem(), nil))
				}
			}
			if !fieldValV.IsNil() {
				fieldElm := fieldValV.Elem() // acquire value referenced by pointer

				if fieldElm.Type().Kind() == reflect.Struct {
					doResetNil(fieldValV, g)
				}
			}
		}
	}
}

func getStructTag(f reflect.StructField, tagName string) string {
	return f.Tag.Get(tagName)
}
