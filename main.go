// Thermostat app main module
//
// Responsible for interfacing with stdin/stdout, parsing program input, and
// interfacing with the owner and owner_collection modules.

package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func readInput() ([]*Owner, [][]string) {
	owners := make([]*Owner, 0)
	queries := make([][]string, 0)

	// Parse data section containing thermostat owners until empty line reached
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		o, err := ParseOwner(line)
		if err != nil {
			log.Fatal(err.Error())
			return nil, nil
		}
		owners = append(owners, o)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err.Error())
		return nil, nil
	}

	// Parse query section containing name-region pairs until end
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" { // check just in case but EOF is expected
			break
		}
		reader := csv.NewReader(strings.NewReader(line))
		reader.Comma = ' ' // set delimiter to space
		fields, err := reader.Read()
		if err != nil {
			log.Fatal(err.Error())
			return nil, nil
		}

		queries = append(queries, fields)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err.Error())
		return nil, nil
	}

	return owners, queries
}

func main() {
	// Read input from stdin and parse into owner and query data structures
	owners, queries := readInput()
	if owners == nil || queries == nil {
		log.Fatal("Loading input data failed.  Exiting.")
		return
	}

	// Sort owners once so O(N^2) cost isn't recurring
	SortOwners(owners)

	// Run each query against the sorted owners slice
	for _, query := range queries {
		result := GetOwnerPercentile(owners, query[0], query[1])
		fmt.Printf("\"%s\" \"%s\" %d\n", query[0], query[1], result)
	}
}
