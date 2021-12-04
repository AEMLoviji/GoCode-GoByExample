package main

import (
	"fmt"
)

func main() {
	items := []int{
		18, 20, 14, 17, 13,
		16, 11, 15, 12, 19,
	}

	//order := 1

	end := len(items)

	for i := 0; i < end; i++ {
		for j := 0; j < (end - 1 - i); j++ {
			// for ascending
			if items[j] > items[j+1] {
				// for descending
				//if items[j] < items[j+1] {
				// dynamic way - does not work
				// if ((items[j] - items[j+1]) * order) > 1 {
				items[j], items[j+1] = items[j+1], items[j]
			}
		}
	}

	fmt.Printf("%v\n", items)
}
