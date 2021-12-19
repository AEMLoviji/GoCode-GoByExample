package main

import "fmt"

func main() {
	square := func(i int) int {
		return i * i
	}

	cubes := func(i int) int {
		return i * i * i
	}

	fmt.Printf("%v\n", calc(square, 3))
	fmt.Printf("%v\n", calc(cubes, 3))

	fmt.Printf("%v\n", calc(func(i int) int {
		return i * i
	}, 3))
}

func calc(f func(int) int, x int) int {
	return f(x)
}
