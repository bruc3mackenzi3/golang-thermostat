package main

import "testing"

func TestParseOwner(t *testing.T) {
	input := `"John Doe" "Canada/Ontario/Toronto" 1.5`
	expected := Owner{
		name:     "John Doe",
		location: "Canada/Ontario/Toronto",
		rValue:   1.5,
	}

	result, _ := ParseOwner(input)
	if *result != expected {
		t.Errorf("Echo was incorrect,") // got: %s, expected: %s.", result, expected)
	}
}
