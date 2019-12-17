package main

import "testing"

func TestParseOwner(t *testing.T) {
	input := `"John Doe" "Canada/Ontario/Toronto" 1.5`
	expected := Owner{
		name:     "John Doe",
		location: "Canada/Ontario/Toronto",
		rValue:   1.5,
	}

	result, err := ParseOwner(input)
	if *result != expected || err != nil {
		t.Errorf("Owner was incorrect")
	}

	// Test erroneous cases
	result, err = ParseOwner(`John Doe Canada/Ontario/Toronto 1.5`)
	if result != nil || err == nil {
		t.Errorf("Error not returned,")
	}

	result, err = ParseOwner(`"John Doe" "Canada/Ontario/Toronto" 1.5 extradata`)
	if result != nil || err == nil {
		t.Errorf("Error not returned,")
	}

	result, err = ParseOwner(`"John Doe" "Canada Ontario/Toronto" 1.5`)
	if result != nil || err == nil {
		t.Errorf("Error not returned,")
	}

	result, err = ParseOwner(`"John Doe" "Canada/Toronto" 1.5`)
	if result != nil || err == nil {
		t.Errorf("Error not returned,")
	}
}
