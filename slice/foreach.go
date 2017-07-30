package slice

import (
	"reflect"
)

func ForEach(slice, feFunc interface{}) {
	sliceVal, fVal := reflect.ValueOf(slice), reflect.ValueOf(feFunc)
	for i := 0; i < sliceVal.Len(); i++ {
		fVal.Call([]reflect.Value{sliceVal.Index(i)})
	}
}
