package main

import "fmt"

// ASSIGNMENT: Comparing structs - considering the following struct write
// a simple program to compare two variables of this type. Write two versions
// of the compare() function: call by value and call by ref.

type book struct {
	name   string
	author string
}

func main() {
	b1 := book{"Go Programming", "Author 1"}
	b2 := book{"Go Programming", "Author 1"}
	b3 := book{"C++ Programming", "Author 2"}

	fmt.Printf("LOC1 - %v\n", b1)

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	fmt.Println(compare1(b1, b2))
	fmt.Println(compare1(b1, b3))
	fmt.Printf("LOC2 - %v\n\n", b1)

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	fmt.Println(compare2(&b1, &b2))
	fmt.Printf("LOC3 - %v\n\n", b1)

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	p1 := &book{"Go Web Programming", "Author 3"}

	p2 := new(book)
	*p2 = book{"Go Web Programming", "Author 3"}

	fmt.Printf("LOC4 - &p1=%v *p1=%v\n", &p1, *p1)
	fmt.Println(compare2(p1, p2))
	fmt.Printf("LOC5 - &p1=%v *p1=%v\n", &p1, *p1)
}

func compare1(p book, q book) bool { // call by value
	if p.name == q.name && p.author == q.author {
		p.name = "Java Programming" // does not change the original variable. Because it is just a copy
		return true
	}
	return false
}

func compare2(p *book, q *book) bool { // call by ref
	if p.name == q.name && p.author == q.author {
		p.name = "Python Programming" // changes the original variable
		return true
	}
	return false
}
