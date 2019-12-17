package main

import (
	"encoding/csv"
	"errors"
	"strconv"
	"strings"
)

type Owner struct {
	name     string
	location string
	rValue   float64
}

// Parses the owner represented in the string input and returns a pointer to an
// Owner struct.  Returns error if input is not.
// The expected format for input is a space-delimited CSV as such:
// "John Doe" "Canada/Ontario/Toronto" 1.5
func ParseOwner(input string) (*Owner, error) {
	// Input data is CSV format with spaces for delimiter so parse using csv package
	reader := csv.NewReader(strings.NewReader(input))
	reader.Comma = ' ' // set delimiter to space
	fields, err := reader.Read()
	if err != nil {
		return nil, err
	}
	if len(fields) != 3 {
		return nil, errors.New("Input data invalid: " + input)
	}

	// Parse rValue as float
	rValue, err := strconv.ParseFloat(fields[2], 64)
	if err != nil {
		return nil, err
	}

	// Validate location
	geo_fields := strings.Split(fields[1], "/")
	if len(geo_fields) != 3 || len(geo_fields[0]) == 0 || len(geo_fields[1]) == 0 || len(geo_fields[2]) == 0 {
		return nil, errors.New("Location field in input string invalid: " + input)
	}

	owner := Owner{
		name:     fields[0],
		location: fields[1],
		rValue:   rValue,
	}
	return &owner, nil
}
