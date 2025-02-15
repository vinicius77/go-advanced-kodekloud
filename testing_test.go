package main

import "testing"

// go test .
func TestCheckEven(t *testing.T) {
	i := 10
	expected := "Yes"
	res := checkEven(i)

	if expected != res {
		t.Errorf("expected: %v, got: %v", expected, res)
	}
}
