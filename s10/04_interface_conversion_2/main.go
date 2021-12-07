package main

import (
	"fmt"
	"strconv"
)

func main() {

	iResult, _ := strconv.Atoi("42")
	fmt.Printf("%d %T \n", iResult, iResult)

	sResult := strconv.Itoa(12)
	fmt.Printf("%s %T \n", sResult, sResult)

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	b, _ := strconv.ParseBool("true")
	f, _ := strconv.ParseFloat("3.1415", 64)
	i, _ := strconv.ParseInt("-42", 10, 64)
	u, _ := strconv.ParseUint("42", 10, 64)

	fmt.Println(b, f, i, u)
	fmt.Printf("%T %T %T %T\n", b, f, i, u)

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	sb := strconv.FormatBool(true)
	sf := strconv.FormatFloat(3.1415, 'E', -1, 64)
	si := strconv.FormatInt(-42, 10)
	su := strconv.FormatUint(42, 10)

	fmt.Println(sb, sf, si, su)
	fmt.Printf("%T %T %T %T\n", sb, sf, si, su)
}
