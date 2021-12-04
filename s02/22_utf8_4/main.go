package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	c := "Hello書"                          // 6*1byte + 1*3byte = 9bytes
	fmt.Println(len(c))                    // number of bytes
	fmt.Println(utf8.RuneCountInString(c)) // actual characters count
}
