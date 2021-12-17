package main

import (
	"fmt"
)

var wordSet1 = []string{"small", "medium", "large"}
var wordSet2 = []string{"red", "green", "blue", "yellow"}

func main() {
	c := make(chan string)

	go queue1(c)
	go queue2(c)

	// fatal error: all goroutines are asleep - deadlock!
	// we need to somehow close channel in occurate way in correct place
	// with semaphore we can achive it
	for val := range c {
		fmt.Println(val) // similar to <-c
	}
}

func queue1(c chan string) {
	for _, w := range wordSet1 {
		c <- w
	}
	//close(c) // unpredicated behaviour!
}

func queue2(c chan string) {
	for _, w := range wordSet2 {
		c <- w
	}
	//close(c) // unpredicated behaviour!
}
