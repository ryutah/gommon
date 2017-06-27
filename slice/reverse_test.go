package slice_test

import (
	"reflect"
	"testing"

	"github.com/ryutah/gommon/slice"
)

func TestReverse(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	got := slice.Reverse(arr).([]int)
	if want := []int{5, 4, 3, 2, 1}; !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}
