package slice_test

import (
	"testing"

	"github.com/ryutah/gommon/slice"
)

func TestIndex(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	got, ok := slice.Index(arr, 3)
	if !ok {
		t.FailNow()
	}
	if want := 2; want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestIndexNoData(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	_, ok := slice.Index(arr, 6)
	if ok {
		t.FailNow()
	}
}
