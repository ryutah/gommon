package slice_test

import (
	"testing"

	"github.com/ryutah/gommon/slice"
)

func TestFind(t *testing.T) {
	s := []string{"foo", "bar", "foobar"}
	i, ok := slice.Find(s, "bar")
	if !ok {
		t.Fail()
	}
	if want := 1; i != i {
		t.Errorf("want %v, got %v", want, i)
	}
}
