package main

import "fmt"

func main() {
	var b1 = true
	var b2 = false

	fmt.Println(b1 || b2)
	fmt.Println(b1 && b2)
	fmt.Println(b1 && !b2)
	fmt.Println(true && !!false)
}
