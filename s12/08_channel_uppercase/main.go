package main

import (
	"fmt"
	"strings"
	"time"
)

// Assignment - Write a program that contains two channels that
// accept texts an convert them to uppercase and lowercase.

func main() {
	text := "Please Modify Me!"

	cu := make(chan string) // cu for cUppercase
	cl := make(chan string) // cu for cLowercase

	fmt.Println("(m) before uppercase()")
	go uppercase(text, cu)
	fmt.Println("(m) before lowercase()")
	go lowercase(text, cl)
	fmt.Println("(m) before both")

	sUpper, sLower := <-cu, <-cl
	fmt.Printf("sUpper=%s sUpper=%s\n", sUpper, sLower)
}

func uppercase(s string, c chan string) {
	fmt.Println("(f) entering uppercase()")
	time.Sleep(1 * time.Second)

	c <- strings.ToUpper(s)
	fmt.Println("(f) exiting from uppercase(); This line may or may not be reached.")
}

func lowercase(s string, c chan string) {
	fmt.Println("(f) entering lowercase()")
	time.Sleep(1 * time.Second)

	c <- strings.ToLower(s)
	fmt.Println("(f) exiting from lowercase(); This line may or may not be reached.")
}
