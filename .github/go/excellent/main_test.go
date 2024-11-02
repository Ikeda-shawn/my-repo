package main

import "testing"

func TestEvenorodd(t *testing.T)  {
	result := Evenorodd(10)
	if result {
		t.Errorf("expected: even, actual: %s", result)
	}
	
}
