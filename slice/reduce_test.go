package slice_test

import (
	"testing"

	"github.com/ryutah/gommon/slice"
)

func TestReduce(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	got := slice.Reduce(arr, 0, func(a, b int) int {
		return a + b
	})
	if want := 15; want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}
