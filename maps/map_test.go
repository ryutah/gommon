package maps_test

import (
	"reflect"
	"testing"

	"github.com/ryutah/gommon/maps"
)

func TestMap(t *testing.T) {
	m := map[string]string{"hoge": "HOGE", "fuga": "FUGA"}
	f := func(k, v string) string {
		return k + v
	}

	want := []string{"hogeHOGE", "fugaFUGA"}
	got, ok := maps.Map(m, f).([]string)

	if !ok {
		t.Fatalf("should return type of []string")
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}
