package main

import "fmt"

// Assignment: Use 'slice' and variadic functions to implement a
// simple stack with three operations:
// push: adding element to the top of stack
// pop: removing one element from the top of stack
// top retrieving the head element of the stack
func main() {
	s := []int{}

	s = push(s, 1, 2, 3)
	fmt.Println(s)

	s = pop(s)
	fmt.Println(s)

	s = push(s, 10, 11)
	fmt.Println(s)

	fmt.Println(top(s))
}

func push(s []int, news ...int) []int {
	return append(s, news...)
}

func pop(s []int) []int {
	return s[:len(s)-1]
}

func top(s []int) int {
	return s[len(s)-1]
}
