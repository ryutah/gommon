package slice

import (
	"reflect"
)

func Index(slice interface{}, v interface{}) (int, bool) {
	sliceVal := reflect.ValueOf(slice)
	for i := 0; i < sliceVal.Len(); i++ {
		val := sliceVal.Index(i).Interface()
		if v == val {
			return i, true
		}
	}
	return -1, false
}
