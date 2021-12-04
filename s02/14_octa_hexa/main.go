package main

import "fmt"

func main() {

	d := 125 // decimal
	fmt.Printf("%8d %#8o %#8x \t %08b\n", d, d, d, d)

	oc := 0175 // octa
	fmt.Printf("%8d %#8o %#8x \t %08b\n", oc, oc, oc, oc)

	x := 0x7d // hexa
	fmt.Printf("%8d %#8o %#8x \t %08b\n", x, x, x, x)

	b := 00000001 // binary
	fmt.Printf("%8d %#8o %#8x \t %08b\n", b, b, b, b)
}
