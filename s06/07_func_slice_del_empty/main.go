package main

import "fmt"

// Assignment: Write a function to accept a slice of string, remove blank strings
// and return the final slice.
func main() {
	data := []string{"Daisy", "Rose", "", "Tulip"}

	//fmt.Printf("%q\n", trimSlice(data)) // form 1
	fmt.Printf("%q\n", trimSlice(data...)) // form 2

	fmt.Printf("%q\n", data)
}

//func trimSlice(data []string) []string { // form 1
func trimSlice(data ...string) []string { // form 2
	var newData []string

	for _, d := range data {
		if d != "" {
			newData = append(newData, d)
		}
	}

	return newData
}
