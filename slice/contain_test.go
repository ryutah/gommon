package slice_test

import (
	"testing"

	"github.com/ryutah/gommon/slice"
)

func TestContain(t *testing.T) {
	s := []string{"foo", "bar", "foobar"}
	if !slice.Contain(s, "bar") {
		t.FailNow()
	}
}

func TestContainNotContained(t *testing.T) {
	s := []string{"foo", "bar", "foobar"}
	if slice.Contain(s, "no") {
		t.Fail()
	}
}

func TestContainsBy(t *testing.T) {
	type Foo struct {
		Foo string
	}
	s := []*Foo{
		{Foo: "foo"},
		{Foo: "bar"},
		{Foo: "foobar"},
	}
	f := &Foo{Foo: "bar"}
	ok := slice.ContainsBy(s, f, func(foo *Foo) bool {
		return foo.Foo == f.Foo
	})
	if !ok {
		t.Fail()
	}
}

func TestContainsByNotContain(t *testing.T) {
	type Foo struct {
		Foo string
	}
	s := []*Foo{
		{Foo: "foo"},
		{Foo: "bar"},
		{Foo: "foobar"},
	}
	f := &Foo{Foo: "BAR"}
	ok := slice.ContainsBy(s, f, func(foo *Foo) bool {
		return foo.Foo == f.Foo
	})
	if ok {
		t.Fail()
	}
}
