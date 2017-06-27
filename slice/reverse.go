package slice

import (
	"reflect"
)

func Reverse(slice interface{}) interface{} {
	sliceType, sliceVal := reflect.TypeOf(slice), reflect.ValueOf(slice)
	retSlice := reflect.MakeSlice(sliceType, sliceVal.Len(), sliceVal.Cap())
	for i := 0; i < sliceVal.Len(); i++ {
		j := (sliceVal.Len() - 1) - i
		retSlice.Index(j).Set(sliceVal.Index(i))
	}
	return retSlice.Interface()
}
