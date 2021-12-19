package main

import "fmt"

func main() {
	nextPos1 := getPositiveInt()

	fmt.Println(nextPos1())
	fmt.Println(nextPos1())
	fmt.Println(nextPos1())

	fmt.Println()
	nextPos2 := getPositiveInt()
	fmt.Println(nextPos2())

	fmt.Printf("'%T' '%T' \n", nextPos1, nextPos1())
	fmt.Printf("%x %x \n", &nextPos1, &nextPos2)
}

func getPositiveInt() func() int {
	i := 0
	fmt.Println("You called getPositiveInt()")

	return func() int {
		fmt.Println("You called inner func()")
		i++
		return i
	}
}
