package main

import "fmt"

// Assignment - implement (k:i, v:i*i) using map.
// example - 2:4, 3:9, 4:16, ...
func main() {
	power2 := make(map[int]int)

	power2[2] = 4
	power2[3] = 9

	fmt.Println(power2)
	fmt.Println(power2[2])
	fmt.Println(power2[3])
	fmt.Println(power2[1]) // 0 - as it does not exist

	delete(power2, 1)
	delete(power2, 3)
	fmt.Println(power2)

	power2[4] = 16
	power2[3] = 9
	power2[5] = 25
	fmt.Println(power2)
	fmt.Println("len is:", len(power2))

	for num, valOfPower2 := range power2 {
		fmt.Printf("power of %d eq to %d\n", num, valOfPower2)
	}

	if val, ok := power2[5]; ok {
		fmt.Printf("Power exist for number 5 and it is eq to %d", val)
	}
}
