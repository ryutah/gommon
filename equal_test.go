package gommon_test

import (
	"testing"

	"github.com/ryutah/gommon"
)

func TestEqual(t *testing.T) {
	a, b := "hogehoge", "hogehoge"
	ok := gommon.Equal(a, b, func(c, d string) bool {
		return c == d
	})
	if !ok {
		t.Fail()
	}
}
