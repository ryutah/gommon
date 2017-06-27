package slice

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func Shuffle(slice interface{}) {
	if sliceType := reflect.TypeOf(slice); sliceType.Kind() != reflect.Slice {
		panic(fmt.Sprintf("%v is not slice", slice))
	}

	sliceVal := reflect.ValueOf(slice)
	for i := 0; i < sliceVal.Len(); i++ {
		rand.Seed(time.Now().UnixNano())
		j := rand.Intn(i + 1)
		buf := sliceVal.Index(i).Interface()
		sliceVal.Index(i).Set(sliceVal.Index(j))
		sliceVal.Index(j).Set(reflect.ValueOf(buf))
	}
}
