package main

import "fmt"

func main() {
	const PI = 3.14159

	radius := 3.0

	area := radius * radius * PI

	fmt.Printf("The area for the circle of radius %5.2f is %5.2f.\n",
		radius, area)
}
