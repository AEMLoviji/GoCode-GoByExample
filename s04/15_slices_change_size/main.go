package main

import "fmt"

func main() {
	var x []int

	fmt.Println(x, len(x), cap(x))

	a := []int{100, 200, 300}

	if aLen := len(a); len(x) <= aLen {
		xLen := aLen
		xCap := xLen + 1

		x = make([]int, xLen, xCap)
		copy(x, a)
	}

	fmt.Println(a, len(a), cap(a))
	fmt.Println(x, len(x), cap(x))
}
