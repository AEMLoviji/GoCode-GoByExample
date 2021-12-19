package main

import "fmt"

func main() {
	//3*f(2)
	//  2*f(1)
	//    1*f(0)
	//	  1
	//	1*1->1
	//  2*1->2
	//3*2->6
	fmt.Println("=>", factorial(3))
	fmt.Println("=>", factorial(4))
	fmt.Println("=>", factorial(7))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	fmt.Print(n, " ")
	return n * factorial(n-1)
}
