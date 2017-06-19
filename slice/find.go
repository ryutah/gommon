package slice

import (
	"fmt"
	"reflect"
)

func Find(slice interface{}, val interface{}) (int, bool) {
	return FindBy(slice, val, func(v interface{}) bool {
		return v == val
	})
}

func FindBy(slice interface{}, val interface{}, compareFunc interface{}) (int, bool) {
	sliceType := reflect.TypeOf(slice)
	if sliceType.Kind() != reflect.Slice {
		panic(fmt.Sprintf("%v is not slice", slice))
	}
	s, f := reflect.ValueOf(slice), reflect.ValueOf(compareFunc)
	for i := 0; i < s.Len(); i++ {
		retValues := f.Call([]reflect.Value{s.Index(i)})
		eql, ok := retValues[0].Interface().(bool)
		if !ok {
			panic("compare function is not return bool")
		}
		if eql {
			return i, true
		}
	}
	return 0, false
}
