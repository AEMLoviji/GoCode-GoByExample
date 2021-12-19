package main

import "fmt"

func main() {
	var x interface{}

	x = mult

	switch i := x.(type) {
	case int:
		fmt.Printf("%T %v", i, i)
	case float64:
		fmt.Printf("%T %v", i, i)
	case bool, string:
		fmt.Println("type is bool or string")
	case func(int) float64:
		fmt.Printf("%T", i)
	case nil:
		fmt.Println("x is nil")
	default:
		fmt.Println("do not know the type")
	}

}

func mult(i int) float64 {
	return float64(i) * 1.5
}
