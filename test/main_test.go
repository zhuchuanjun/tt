// main_test.go
package main

import "testing"

func TestAdd(t *testing.T) {
	if Add(2, 3) != 5 {
		t.Error("Expected 2 + 3 to equal 5")
	}
}

func TestSubtract(t *testing.T) {
	if Subtract(5, 3) != 2 {
		t.Error("Expected 5 - 3 to equal 2")
	}
}
