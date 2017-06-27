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

	slice.Shuffle(arr)

	for _, av := range arr {
		for i := 0; i < len(want); i++ {
			if av == want[i] {
				want = append(want[:i], want[i+1:]...)
			}
		}
	}

	if reflect.DeepEqual(arr, want) {
		t.Fatal("array did not be shuffled")
	}

	if len(want) != 0 {
		t.Errorf("remains %v, array : %v", want, arr)
	}
}
