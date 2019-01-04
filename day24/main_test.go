package main

import "testing"

func TestReadInput(t *testing.T) {
	immune, infect := readInput("input-test.txt")
	if l := len(immune); l != 2 {
		t.Fatalf("Expected %d immune groups but %d found", 2, l)
	}

	if l := len(infect); l != 2 {
		t.Fatalf("Expected %d infect groups but %d found", 2, l)
	}
}
