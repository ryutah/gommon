package slice

import (
	"reflect"
)

func FindAll(slice interface{}, compareFunc interface{}) (interface{}, bool) {
	// XXX 新sliceの生成、値のコピー方法がいけてない。
	// reflectを使用したsliceの操作などを少し詳しく調べた方が良さげ
	sliceType := reflect.TypeOf(slice)
	sliceVal, fVal := reflect.ValueOf(slice), reflect.ValueOf(compareFunc)
	retSlice := reflect.MakeSlice(sliceType, 0, sliceVal.Cap())

	for i := 0; i < sliceVal.Len(); i++ {
		val := sliceVal.Index(i)
		retValues := fVal.Call([]reflect.Value{val})
		eql, ok := retValues[0].Interface().(bool)
		if !ok {
			panic("compare function is not return bool")
		}
		if eql {
			// XXX 下のような感じが理想だけど、SetLen `reflect.value.setlen using unaddressable` というエラーがは発生する
			//   retSlice.SetLen(retSlice.Len() + 1)
			//   reflect.Append(retSlice, val)
			// どうも、reflect.MakeSliceでsliceを生成のはダメらしい
			// 参考
			//   https://stackoverflow.com/questions/25384640/why-golang-reflect-makeslice-returns-un-addressable-value
			//   https://groups.google.com/forum/#!topic/golang-nuts/gzuURvuCmMc
			newSlice := reflect.MakeSlice(sliceType, retSlice.Len()+1, retSlice.Cap())
			for i := 0; i < retSlice.Len(); i++ {
				newSlice.Index(i).Set(retSlice.Index(i))
			}
			newSlice.Index(retSlice.Len()).Set(val)
			retSlice = newSlice
		}
	}

	return retSlice.Interface(), retSlice.Len() >= 1
}
