package gommon

import (
	"reflect"
)

func Equal(a, b interface{}, compareFunc interface{}) bool {
	var (
		va = reflect.ValueOf(a)
		vb = reflect.ValueOf(b)
		vf = reflect.ValueOf(compareFunc)
	)
	retValues := vf.Call([]reflect.Value{va, vb})
	eql, ok := retValues[0].Interface().(bool)
	if !ok {
		panic("compare function return value is not bool")
	}
	return eql
}
