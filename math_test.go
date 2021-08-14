package main

import "testing"

func TestSum(t *testing.T) {
	expected := 48
	actual := Sum(15, 33)

	if actual != expected {
		t.Errorf("Invalid result: Expected = %d Actual = %d", expected, actual)
	}
}