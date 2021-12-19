package main

import (
	"fmt"
	"strings"
	//tjarratt "github.com/tjarratt/babble"
)

// Assignment:Write a pipeline including two goroutines:
// 1) The 'newWords' goroutine that produces random words and send them to a channel;
// 2) The 'uWords' (uppercase words) goroutine that converts the randomly-generated
// words to uppercase;
// Use anonymous functions to write your goroutines.
//

func main() {
	//babbler := tjarratt.NewBabbler()
	newWords := make(chan string)
	uWords := make(chan string)

	go func() { // 'newWords' goroutine
		for i := 0; i < 5; i++ {
			newWords <- fmt.Sprintf("word.%d", i) //babbler.Babble()
		}
		close(newWords)
	}()

	go func() { // 'uWords' goroutine
		for w := range newWords {
			uWords <- w + " --> " + strings.ToUpper(w)
		}
		close(uWords)
	}()

	for w := range uWords {
		fmt.Println(w)
	}

}
