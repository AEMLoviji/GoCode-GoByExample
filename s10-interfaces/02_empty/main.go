package main

import "fmt"

type athlete struct {
	name    string
	country string
}

type football struct { // means footbal player
	athlete
	position string
}

type tennis struct {
	athlete
	rightHanded bool
}

type athletes interface{}

func playerInfo(a interface{}) {
	fmt.Println(a)
}

func main() {
	messi := football{}
	pele := football{}
	federer := tennis{}
	nadal := tennis{}

	favAthletes := []athletes{messi, pele, federer, nadal}

	for k, v := range favAthletes {
		fmt.Println(k, " ‐ ", v)
	}

	messi = football{athlete{"Leo Messi", "Argentina"}, "Attacker"}
	federer = tennis{athlete{"Roger Federer", "Switzerland"}, true}
	playerInfo(messi)
	playerInfo(federer)

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	pele = football{athlete{"Pele", "Brazil"}, "Attacker"}
	nadal = tennis{athlete{"Rafael Nadal", "Spain"}, false}
	favAthletes2 := []interface{}{messi, pele, federer, nadal}
	fmt.Println(favAthletes2)
}
