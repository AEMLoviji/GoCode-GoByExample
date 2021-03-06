package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string, 3)
	fmt.Printf("capacity:%d length:%d \n", cap(c), len(c))

	c <- "Message 1"
	c <- "Message 2"
	fmt.Printf("capacity:%d length:%d \n", cap(c), len(c))

	c <- "Message 3"
	fmt.Printf("capacity:%d length:%d \n", cap(c), len(c))

	time.Sleep(1 * time.Second)
	fmt.Println(<-c)
	time.Sleep(1 * time.Second)
	fmt.Println(<-c)

	fmt.Println(<-c) //ok
	//fmt.Println(<-c) //fatal error: all goroutines are asleep - deadlock!

	fmt.Printf("capacity:%d length:%d \n", cap(c), len(c))
}
