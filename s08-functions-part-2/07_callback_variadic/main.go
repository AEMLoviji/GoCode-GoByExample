package main

import "fmt"

// Assignment: Write a callback with two parameters:
// 1. A function that accepts a variable number of integers and return an integer.
// In practice this could be a function to add or multiply, ... some numbers.
// 2. The list of integers (with a variable number of elements) to be added, or multiplied, ...
// Use the techniques learned in the last example, as well the ones to handle
// a variable number of parameters.
func main() {
	add := func(nums ...int) int {
		total := 0
		for _, num := range nums {
			total += num
		}
		return total
	}

	fmt.Println(add(1, 2, 3, 4))

	nums := []int{1, 2, 3, 4}
	fmt.Println(add(nums...))

	fmt.Printf("%v\n", calc2(add, nums))

	mult := func(nums ...int) int {
		total := 1
		for _, num := range nums {
			total *= num
		}
		return total
	}

	fmt.Printf("%v\n", calc2(mult, nums))
}

func calc2(f func(...int) int, x []int) int {
	return f(x...)
}
