package main

import "fmt"

func main() {
	x := 1
	p := &x
	q := &x

	fmt.Printf(" x =>  type=%T value=%d \n", x, x)
	fmt.Printf("&x =>  type=%T value=%d \n", &x, &x)

	fmt.Println()
	fmt.Printf("*p =>  type=%T value=%d \n", *p, *p)
	fmt.Printf(" p =>  type=%T value=%d \n", p, p)
	fmt.Printf("&p =>  type=%T value=%d \n", &p, &p)

	fmt.Println()
	fmt.Printf("*q =>  type=%T value=%d \n", *q, *q)
	fmt.Printf(" q =>  type=%T value=%d \n", q, q)
	fmt.Printf("&q =>  type=%T value=%d \n", &q, &q)
}
