package main

import (
	"fmt"
	"strings"
)

// Assignment: Rewrite the previous assignment (pipeline-like)
// using functions and unidirectional Channels.

func main() {
	newWords := make(chan string)
	uWords := make(chan string)

	go sendWords(newWords)
	go convertWords(uWords, newWords)

	printWords(uWords)
}

func sendWords(out chan<- string) {
	for i := 0; i < 5; i++ {
		out <- fmt.Sprintf("word.%d", i) //babbler.Babble()
	}
	close(out)
}

func convertWords(out chan<- string, in <-chan string) {
	for w := range in {
		out <- w + " --> " + strings.ToUpper(w)
	}
	close(out)
}

func printWords(in <-chan string) {
	for w := range in {
		fmt.Println(w)
	}
}
