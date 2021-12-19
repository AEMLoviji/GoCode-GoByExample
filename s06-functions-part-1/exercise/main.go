package main

import (
	"fmt"
	"time"
)

var delayTime = 100 * time.Millisecond

func main() {
	executionTime := 3 * time.Second
	start := time.Now()

	fmt.Printf("Program will end in about %v.\n", executionTime)
	fmt.Print("Operation in progress ... ")

	s := `\|/-`
	i := 0
	for {
		printSpinningSymbol(string(s[i]))

		if timeEllapsed := time.Since(start); timeEllapsed > executionTime {
			fmt.Println("Done")
			fmt.Println("Elapsed Time: ", timeEllapsed)
			break
		}

		i = (i + 1) % 4
	}
}

func printSpinningSymbol(symbol string) {
	fmt.Print(symbol)
	time.Sleep(delayTime)
	fmt.Print("\b")

}
