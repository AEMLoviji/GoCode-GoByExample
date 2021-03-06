package main

import "fmt"

func main() {
	type myType float64

	var total myType

	total = 44
	fmt.Printf("%.2f %T \n", total, total)

	var total2 float64

	// total2 = total // compiler error- can not use myType as float64
	total2 = float64(total)
	fmt.Printf("%.2f %T \n", total2, total2)

	//if total == total2 { // compiler error - mismatched types
	//	fmt.Println("x")
	//}
}
