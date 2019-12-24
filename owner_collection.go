package main

import "sort"

// SortOwners Sorts the slice of owners by the rValue field, in ascending order.
// Note the sort is stable (meaning order of equal items is preserved) and the
// sort is done in place, hence nothing is returned.
// Time complexity: O(N^2)
// Space complexity: O(N)
func SortOwners(owners []*Owner) {
	sort.SliceStable(owners, func(i, j int) bool {
		return owners[i].rValue < owners[j].rValue
	})
}

// GetOwnerPercentile - Given an owner name search the Owner slice for their
// entry and return their ranking rValue
func GetOwnerPercentile(owners []*Owner, name string, location string) int {
	prefixLen := len(location)
	count := 0
	ownerIndex := -1
	for _, owner := range owners {
		if location == owner.location[0:prefixLen] {
			count++
			if name == owner.name {
				ownerIndex = count
			}
		}
	}
	return calculatePercentile(ownerIndex, count)
}

// Helper function to calculate the percentile of an item at index i in dataset
// of size.  The percentile is returned as an integer between 1 and 10 meaning
// e.g. 1 represents the item is in the 1st to 10th
func calculatePercentile(index, size int) int {
	var percentileRounded int // needed because set in if/else block

	// First calculate the percentile as a float to preserve the decimal
	percentile := (float32(index) * 10.0) / float32(size)

	// Drop decimal by casting to int then back to float, then compare to test
	// for presence of decimal.  If so add one to round up to next integer.
	if float32(int(percentile)) != percentile {
		percentileRounded = int(percentile + 1.0)
	} else {
		percentileRounded = int(percentile)
	}

	return percentileRounded
}
