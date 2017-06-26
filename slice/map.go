package slice

import (
	"fmt"
	"reflect"
)

func Map(slice interface{}, mapFunc interface{}) interface{} {
	sliceType := reflect.TypeOf(slice)
	if sliceType.Kind() != reflect.Slice {
		panic(fmt.Sprintf("%v is not slice", slice))
	}
	sliceVal, fVal := reflect.ValueOf(slice), reflect.ValueOf(mapFunc)

	fReturnType := reflect.TypeOf(mapFunc).Out(0)
	rSliceType := reflect.SliceOf(fReturnType)
	rSlice := reflect.MakeSlice(rSliceType, sliceVal.Len(), sliceVal.Cap())

	for i := 0; i < sliceVal.Len(); i++ {
		v := sliceVal.Index(i)
		retVal := fVal.Call([]reflect.Value{v})
		rSlice.Index(0).Set(retVal[0])
	}
	return rSlice.Interface()
}
