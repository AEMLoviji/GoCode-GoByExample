package main

import (
	"fmt"
	"sort"
)

type person []string

func (a person) Len() int           { return len(a) }
func (a person) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a person) Less(i, j int) bool { return a[i] < a[j] }

func main() {
	s := person{"Susan", "Tyler", "Ava", "Nick"}
	fmt.Println(s)

	sort.Sort(s)
	fmt.Println(s)

	sort.Sort(sort.Reverse(s))
	fmt.Println(s)
}
