package main

import "testing"

func TestSortOwners(t *testing.T) {
	owners := []*Owner{
		&Owner{
			name:     "Alicia Yazzie",
			location: "US/Arizona/Phoenix",
			rValue:   5.532,
		},
		&Owner{
			name:     "John Doe",
			location: "Canada/Ontario/Toronto",
			rValue:   1.5,
		},
	}
	expected := []*Owner{owners[1], owners[0]}

	// Test sort where input slice is unsorted.  Expect items reversed
	SortOwners(owners)
	if owners[0] != expected[0] || owners[0].rValue > owners[1].rValue {
		t.Errorf("SortOwner failed, rValues not in ascending order: %f %f", owners[0].rValue, owners[1].rValue)
	}

	// Test stable sort with equal items not changing place
	owners[1].rValue = 1.5
	SortOwners(owners)
	if owners[0] != expected[0] || owners[0].rValue > owners[1].rValue {
		t.Errorf("SortOwner failed, rValues not in ascending order: %f %f", owners[0].rValue, owners[1].rValue)
	}
}

func TestCalculatePercentile(t *testing.T) {
	expected := 4
	result := calculatePercentile(2, 5)
	if result != expected {
		t.Errorf("Got %d expected %d", result, expected)
	}

	expected = 2
	result = calculatePercentile(2, 10)
	if result != expected {
		t.Errorf("Got %d expected %d", result, expected)
	}

	expected = 1
	result = calculatePercentile(1, 100)
	if result != expected {
		t.Errorf("Got %d expected %d", result, expected)
	}

	expected = 1
	result = calculatePercentile(10, 100)
	if result != expected {
		t.Errorf("Got %d expected %d", result, expected)
	}

	expected = 2
	result = calculatePercentile(11, 100)
	if result != expected {
		t.Errorf("Got %d expected %d", result, expected)
	}

	expected = 10
	result = calculatePercentile(99, 100)
	if result != expected {
		t.Errorf("Got %d expected %d", result, expected)
	}

	expected = 10
	result = calculatePercentile(100, 100)
	if result != expected {
		t.Errorf("Got %d expected %d", result, expected)
	}
}

func TestGetOwnerPercentile(t *testing.T) {
	owners := []*Owner{
		&Owner{
			name:     "John Doe",
			location: "Canada/Ontario/Toronto",
			rValue:   1.5,
		},
		&Owner{
			name:     "Adam Xin",
			location: "Canada/British Columbia/Vancouver",
			rValue:   2.110,
		},
	}
	expected := 5
	result := GetOwnerPercentile(owners, "John Doe", "Canada")
	if result != expected {
		t.Errorf("Got %d expected %d", result, expected)
	}

	expected = 10
	result = GetOwnerPercentile(owners, "Adam Xin", "Canada")
	if result != expected {
		t.Errorf("Got %d expected %d", result, expected)
	}

	expected = 10
	result = GetOwnerPercentile(owners, "John Doe", "Canada/Ontario")
	if result != expected {
		t.Errorf("Got %d expected %d", result, expected)
	}
}
