package main

import "fmt"

func main() {
	isSunny := true

	if isSunny {
		fmt.Println("Enjoy your day!")
	} else {
		fmt.Println("Tomorrow is another day")
	}

	isSaturday := true

	if isSunny && isSaturday {
		fmt.Println("Let's swim outdoor.")
	}

}
