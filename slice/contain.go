package slice

import (
	"fmt"
	"reflect"
)

func Contain(slice interface{}, val interface{}) bool {
	return ContainsBy(slice, val, func(v interface{}) bool {
		return v == val
	})
}

func ContainsBy(slice interface{}, val interface{}, funcCompare interface{}) bool {
	sliceType := reflect.TypeOf(slice)
	if sliceType.Kind() != reflect.Slice {
		panic(fmt.Sprintf("%v is not slice", slice))
	}

	s, f := reflect.ValueOf(slice), reflect.ValueOf(funcCompare)
	for i := 0; i < s.Len(); i++ {
		retValues := f.Call([]reflect.Value{s.Index(i)})
		isContain, ok := retValues[0].Interface().(bool)
		if !ok {
			panic("compare function is not return bool")
		}
		if isContain {
			return true
		}
	}
	return false
}
