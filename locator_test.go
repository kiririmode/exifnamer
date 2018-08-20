package main

import (
	"reflect"
	"testing"
)

func TestNewFixedLocator(t *testing.T) {
	var l Locator = NewFixedLocator("hoge")

	actual := reflect.TypeOf(l).String()
	expected := "*main.FixedLocator"
	if actual != expected {
		t.Errorf("NewFixedLocator doesn't return NewFixedLocator: %T", l)
	}
	if _, ok := l.(Locator); !ok {
		t.Error("An object NewFixedLocator returns does not inplements Locator interface")
	}
}

func TestFixedLocator_Dest(t *testing.T) {
	sut := NewFixedLocator("hoge")
	actual, err := sut.Dest("fuga")
	if err != nil {
		t.Errorf("error should be nil but: %s", err)
	}

	expected := "hoge"
	if actual != expected {
		t.Errorf("got: %v, want: %v\n", actual, expected)
	}
}

func TestOriginalFileLocator_Dest(t *testing.T) {
	sut := &OriginalFileLocator{}

	testCases := []struct {
		path     string
		expected string
	}{
		{path: "/a/b/c", expected: "/a/b"},
		{path: "/a/b/c/", expected: "/a/b/c"},
		{path: "/a", expected: "/"},
		{path: "./a", expected: "."},
		{path: "a", expected: "."}}

	for _, tc := range testCases {
		actual, err := sut.Dest(tc.path)
		if err != nil {
			t.Errorf("error should be nil, but: %s", err)
		}
		if actual != tc.expected {
			t.Errorf("expects: %s, but actual: %s", tc.expected, actual)
		}
	}
}
