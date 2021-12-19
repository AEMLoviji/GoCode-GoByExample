package main

import "fmt"

func main() {
	var i = 10
	var f = 10.35
	fmt.Println(f + float64(i))
	fmt.Println(int(f) + i)

	var r rune = 'a'
	var i2 int32 = 'b'
	fmt.Println(r, i2)
	fmt.Println(string(r), string(i2))

	b := []byte{'t', 'e', 's', 't'}
	fmt.Println(b, string(b))

	s := "test"
	fmt.Println([]byte(s), s)
}
