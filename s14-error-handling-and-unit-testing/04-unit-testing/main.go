package main

import (
	"fmt"

	"github.com/aemloviji/GoCode-GoByExample/s14-unit-testing/myutil/myutil"
)

func main() {
	testCase1()
	fmt.Println()
	testCase2()
}

func testCase1() {
	fmt.Printf("(p) Rectangle Area= %d\n", myutil.Area(2, 3))
	fmt.Printf("(p) Rectangle Area= %d\n", myutil.Perim(2, 3))

	fmt.Printf("(p) Rectangle Area= %d\n", myutil.Area(5, 10))
	fmt.Printf("(p) Rectangle Area= %d\n", myutil.Perim(5, 10))

	fmt.Printf("(p) Rectangle Area= %d\n", myutil.Area(4, 2))
	fmt.Printf("(p) Rectangle Area= %d\n", myutil.Perim(4, 2))
}

func testCase2() {
	fmt.Printf("(p) Average= %.2f\n", myutil.Avg([]float32{4, 6}))
	fmt.Printf("(p) Average= %.2f\n", myutil.Avg([]float32{3, -3}))
	fmt.Printf("(p) Average= %.2f\n", myutil.Avg([]float32{3, -3, 5, -5, 4}))
	fmt.Printf("(p) Average= %.2f\n", myutil.Avg([]float32{1, 1, 1, 100}))
}
