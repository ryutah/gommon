package slice

import (
	"reflect"
)

func Reduce(slice interface{}, initVal interface{}, reduceFunc interface{}) interface{} {
	var (
		sliceVal  = reflect.ValueOf(slice)
		reduceVal = reflect.ValueOf(initVal)
		funcVal   = reflect.ValueOf(reduceFunc)
	)

	for i := 0; i < sliceVal.Len(); i++ {
		reduceVal = funcVal.Call([]reflect.Value{reduceVal, sliceVal.Index(i)})[0]
	}
	return reduceVal.Interface()
}
