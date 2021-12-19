package main

import "fmt"

func main() {
	scores1 := []float32{91, 82, 99}
	fmt.Println(scores1)
	fmt.Printf("avarage: %.2f\n", avg(scores1))
	fmt.Println(scores1)

	scores2 := []float32{72, 81, 78, 91, 68}
	fmt.Printf("avarage: %.2f\n", avg(scores2))
}

func avg(scores []float32) float32 {
	var total float32
	for _, score := range scores {
		total += score
	}

	// slice is reference type so we pass it by reference
	scores[0] = 1000

	return total / float32(len(scores))
}
