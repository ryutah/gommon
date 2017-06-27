package slice_test

import (
	"reflect"
	"testing"

	"github.com/ryutah/gommon/slice"
)

func TestShuffle(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	want := make([]int, len(arr))
	copy(want, arr)

	got := slice.Shuffle(arr).([]int)

	if !reflect.DeepEqual(arr, want) {
		t.Errorf("given array must not be changed. source array %v, want %v", arr, want)
	}

	for _, g := range got {
		for i := 0; i < len(want); i++ {
			if g == want[i] {
				want = append(want[:i], want[i+1:]...)
			}
		}
	}

	if reflect.DeepEqual(got, want) {
		t.Fatal("array did not be shuffled")
	}

	if len(want) != 0 {
		t.Errorf("remains %v, array : %v", want, got)
	}
}
