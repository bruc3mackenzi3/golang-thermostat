package main

import "testing"

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
