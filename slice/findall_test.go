package slice_test

import (
	"reflect"
	"testing"

	"github.com/ryutah/gommon/slice"
)

func TestFindAll(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	got, ok := slice.FindAll(arr, func(a int) bool {
		return a%2 == 0
	})
	if !ok {
		t.Fatal()
	}
	if want := []int{2, 4, 6}; !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}
