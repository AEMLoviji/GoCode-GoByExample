package main

import (
	"fmt"
	"strings"

	"github.com/aemloviji/GoCode-GoByExample/s09-12-exported/athletes"
)

func changeAthleteName1(p athletes.Player) {
	p.Name = "Anderson Silva-Spider"
}

func changeAthleteName2(p *athletes.Player) {
	p.Name = "Anderson Silva-Spider"
	p.Country = strings.ToUpper(p.Country)
}

func main() {
	player1 := athletes.Player{"Anderson Silva", "MMA", 43, athletes.Info{"Brazil", "Black"}}
	fmt.Println("(1) player1:", player1)

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	player2 := athletes.Player{
		Name:  "Anderson Silva",
		Sport: "MMA",
		Age:   43,
		Info: athletes.Info{
			Country:   "Brazil",
			HairColor: "Black",
		},
	}

	fmt.Println("(1) player2:", player2)

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	changeAthleteName1(player1)
	fmt.Println("(2) player1:", player1)

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	changeAthleteName2(&player1)
	fmt.Println("(3) player1:", player1)

	fmt.Println("(4) player1:", *player1.ToLowercase())
}
