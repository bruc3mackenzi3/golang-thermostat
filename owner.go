package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"sort"
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
		fmt.Println(err)
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

// Sorts the slice of owners by the rValue field, in ascending order.  Note the
// sort is stable (meaning order of equal items is preserved) and the sort is
// done in place, hence nothing is returned.
// Time complexity: O(N^2)
// Space complexity: O(N)
func SortOwners(owners []*Owner) {
	sort.SliceStable(owners, func(i, j int) bool {
		return owners[i].rValue < owners[j].rValue
	})
}

// Given an owner name search the Owner slice for their entry and return their
// ranking rValue
func GetOwnerPercentile(owners []*Owner, name string, location string) int {
	prefix_len := len(location)
	count := 0
	owner_index := -1
	for _, owner := range owners {
		if location == owner.location[0:prefix_len] {
			count++
			if name == owner.name {
				owner_index = count
			}
		}
	}
	return calculatePercentile(owner_index, count)
}

// Helper function to calculate the percentile of an item at index i in dataset
// of size.  The percentile is returned as an integer between 1 and 10 meaning
// e.g. 1 represents the item is in the 1st to 10th
func calculatePercentile(index, size int) int {
	return (index * 10) / size
}
