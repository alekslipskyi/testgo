package helpers

import (
	"reflect"
)

func OmitPrivateFields(obj interface{}) {
	svPTR := reflect.ValueOf(obj).Elem()
	sv := reflect.Indirect(svPTR)
	st := sv.Type()

	for i := 0; i < sv.NumField(); i++ {
		key := st.Field(i)

		access := key.Tag.Get("access")
		if access == "private" {
			fieldValue := sv.Field(i)

			fieldValue.Set(reflect.Zero(fieldValue.Type()))
		}
	}
}
