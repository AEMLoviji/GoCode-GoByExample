package main

import "fmt"

func main() {
	x := 10

	fmt.Printf("x=%d\n", x)
	f1(x)
	fmt.Printf("x=%d\n", x)
}

func f1(y int) {
	fmt.Printf("(f1) y=%d\n", y)
	y += 2
	fmt.Printf("(f1) y=%d\n", y)
}
