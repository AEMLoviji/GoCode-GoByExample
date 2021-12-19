package main

import "fmt"

func main() {

	defer func() {
		errMsg := recover()
		fmt.Println(errMsg)
	}()

	var x map[string]int
	x["key"] = 10 // panic created by system when we try to assign to entry in nil map
	fmt.Println(x)

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~~ ~ ~ ~ ~ ~
	// panic("BOO!!!") // manually make panic
	// fmt.Println("This line won't be reached!")
}
