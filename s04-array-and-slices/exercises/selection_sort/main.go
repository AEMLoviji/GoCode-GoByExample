package main

import (
	"fmt"
)

func main() {
	items := []int{
		18, 20, 14, 17, 13,
	}

	end := len(items)

	for i := 0; i < end-1; i++ {
		minIdx := i
		for j := i + 1; j < end; j++ {
			if items[minIdx] > items[j] {
				minIdx = j
			}
		}

		if items[minIdx] < items[i] {
			items[minIdx], items[i] = items[i], items[minIdx]
		}
	}

	fmt.Printf("%v\n", items)
}
