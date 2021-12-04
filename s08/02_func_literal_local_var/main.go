package main

import "fmt"

func main() {
	factor := 2

	mult := func(i, j int) int {
		return i * j * factor
	}

	fmt.Println(mult(3, 4))
}
