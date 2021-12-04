package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const minRank = 1
	const maxRank = 13

	const minSuit = 1
	const maxSuit = 4

	rand.Seed(time.Now().UnixNano())
	rankIdx := minRank + rand.Intn(maxRank-minRank)
	suitIdx := minSuit + rand.Intn(maxSuit-minSuit)

	var rank = [maxRank]string{"Ace", "2", "3", "4", "5",
		"6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	var suit = [maxSuit]string{"Clubs", "Diamonds", "Hearts", "Spades"}

	fmt.Printf("You picked %s of %s.\n",
		rank[rankIdx-1], suit[suitIdx-1])
}
