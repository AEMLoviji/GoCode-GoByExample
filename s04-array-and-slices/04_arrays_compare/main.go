package main

import "fmt"

func main() {
	x2 := [...]int{10, 15, 30}
	x3 := [3]int{10, 15, 30}

	fmt.Println(x2 == x3)
}
