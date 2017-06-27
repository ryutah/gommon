package slice

import (
	"math/rand"
	"reflect"
	"time"
)

func Shuffle(slice interface{}) interface{} {
	sliceType, sliceVal := reflect.TypeOf(slice), reflect.ValueOf(slice)
	retSlice := reflect.MakeSlice(sliceType, sliceVal.Len(), sliceVal.Cap())
	reflect.Copy(retSlice, sliceVal)
	for i := 0; i < retSlice.Len(); i++ {
		rand.Seed(time.Now().UnixNano())
		j := rand.Intn(i + 1)
		buf := retSlice.Index(i).Interface()
		retSlice.Index(i).Set(retSlice.Index(j))
		retSlice.Index(j).Set(reflect.ValueOf(buf))
	}
	return retSlice.Interface()
}
