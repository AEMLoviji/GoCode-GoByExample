package main

import (
	"fmt"
	"sort"
)

type SortBy []string

func (a SortBy) Len() int {
	return len(a)
}
func (a SortBy) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a SortBy) Less(i, j int) bool {
	return a[i] < a[j]
}

func main() {

	n := []int{7, 2, 10, 14, 1, 14, 9}
	fmt.Println(n)
	sort.Ints(n)
	fmt.Println(n)

	s := []string{"Susan", "Tyler", "Ava", "Nick"}
	fmt.Println(s)
	sort.Strings(s)
	fmt.Println(s)

}
