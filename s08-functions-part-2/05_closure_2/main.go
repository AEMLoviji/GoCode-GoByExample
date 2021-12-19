package main

import "fmt"

// Assignment: Using 'closure' write an addBy() function to add x to the
// previous value of the function. For instance,
// 		addCounter := addBy(); Println(addCounter(2)); Println(addCounter(-1))
// Also write a multBy() function to multiply by x. For instance,
// 		multCounter := multBy(); Println(multCounter(3)); Println(multCounter(-2))
func main() {
	addCounter, multCounter := addBy(), multBy()

	fmt.Print(addCounter(2), " ")
	fmt.Print(addCounter(3), " ")
	fmt.Print(addCounter(-1), "\n")

	fmt.Print(multCounter(3), " ")
	fmt.Print(multCounter(4), " ")
	fmt.Print(multCounter(-2))
}

func addBy() func(int) int {
	total := 0
	return func(n int) int {
		total += n
		return total
	}
}

func multBy() func(int) int {
	total := 1
	return func(n int) (ret int) {
		total *= n
		ret = total
		return
	}
}
