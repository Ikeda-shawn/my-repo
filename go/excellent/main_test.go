package main

import "testing"

func Evenorodd(t *testing.T) {
	result := Evenorodd(10)
	if result != "even" {
		t.Errof("expected: even, actual: %s", result)
	}
}
