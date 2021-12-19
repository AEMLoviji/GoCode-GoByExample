package main

import "fmt"

func main() {
	team1 := []string{"Messi", "Pele", "Maradonna", goalkeeper()}

	team := append([]string{"Maldini", "Cafu", "Carlos", "Beckenbauer"}, team1...)
	team = append(team, midfielders()...)

	for i, name := range team {
		fmt.Println(i, name)
	}
}

func goalkeeper() string {
	return "Buffoon"
}

func midfielders() []string {
	return []string{"Iniesta", "Zidane", "Platini"}
}
