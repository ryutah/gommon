package slice_test

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/ryutah/gommon/slice"
)

func TestMap(t *testing.T) {
	arr := []string{"1", "2", "3", "4", "5"}
	got := slice.Map(arr, func(s string) int {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(fmt.Sprintf("ERR : %v", err))
		}
		return i
	}).([]int)
	if want := []int{1, 2, 3, 4, 5}; reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}
