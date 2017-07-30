package slice_test

import (
	"testing"

	"github.com/ryutah/gommon/slice"
)

func TestForEach(t *testing.T) {
	s := []string{
		"hoge",
		"fuga",
		"hogefuga",
	}
	called := map[string]bool{
		"hoge":     false,
		"fuga":     false,
		"hogefuga": false,
	}
	f := func(s string) {
		called[s] = true
	}

	slice.ForEach(s, f)

	for k, v := range called {
		if !v {
			t.Errorf("%v should be true", k)
		}
	}
}
