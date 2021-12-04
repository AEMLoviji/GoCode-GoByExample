package main

import "fmt"

func main() {
	var i int
	for i = 0; i < 15; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}

	fmt.Printf("\n\n%d ", fibonacci(4))
}

func fibonacci(n int) int {

	//fmt.Print(n, " ") // for debugging purposes
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fibonacci(n-1) + fibonacci(n-2)
}
