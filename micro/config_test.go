package main

import (
	"reflect"
	"testing"
)

func TestFilterOutNonJsonPaths(t *testing.T) {
	paths := []string{"abc.json", "bad.file", "my.json"}

	valid := validConfigPaths(paths)

	if len(valid) != 2 {
		t.Error("expected 2 valid paths")
	}

	if !reflect.DeepEqual(valid, []string{"abc.json", "my.json"}) {
		t.Error("expected abc.json and my.json")
	}
}
