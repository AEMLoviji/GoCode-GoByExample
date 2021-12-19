package main

import (
	"fmt"
	"sort"
)

// Assignment: Maps are unorders in naturally.
// Write a program to use a "map[string]float64" data map and sort it.
func main() {
	sal := map[string]float64{
		"Blake":  60000.00,
		"Parker": 120000.50,
		"Dakota": 93000.00,
	}

	// create slice of names
	names := make([]string, 0, len(sal))
	for n := range sal {
		names = append(names, n)
	}

	// sort names slice
	sort.Strings(names)

	// iterate over names slice and output result from dict by corresponding key
	for _, n := range names {
		fmt.Printf("%s\t%.2f\n", n, sal[n])
	}
}
