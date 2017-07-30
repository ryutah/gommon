package maps

import (
	"reflect"
)

func Map(m interface{}, mapFunc interface{}) interface{} {
	mVal, fVal := reflect.ValueOf(m), reflect.ValueOf(mapFunc)
	retType := reflect.TypeOf(mapFunc).Out(0)
	retSlice := reflect.MakeSlice(reflect.SliceOf(retType), mVal.Len(), mVal.Len())

	for i, k := range mVal.MapKeys() {
		v := mVal.MapIndex(k)
		rets := fVal.Call([]reflect.Value{k, v})
		retSlice.Index(i).Set(rets[0])
	}

	return retSlice.Interface()
}
