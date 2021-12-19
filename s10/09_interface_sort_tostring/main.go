package main

import (
	"fmt"
	"sort"
)

type player struct {
	Name    string
	Country string
}

func (p player) String() string {
	return fmt.Sprintf("toString() - %s: %s", p.Name, p.Country)
}

type byCountry []player

func (a byCountry) Len() int           { return len(a) }
func (a byCountry) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byCountry) Less(i, j int) bool { return a[i].Country < a[j].Country }

func main() {
	p := []player{
		{"Roger Federer", "Switzerland"},
		{"Lionel Messi", "Argentina"},
		{"Michael Jordan", "USA"},
	}

	fmt.Println(p[0])
	fmt.Println(&p[0])
	fmt.Println(p)

	fmt.Println()
	sort.Sort(byCountry(p))
	fmt.Println(p)
}
