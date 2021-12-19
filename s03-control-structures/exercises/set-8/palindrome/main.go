package main

import "fmt"

func main() {
	var word string
	fmt.Print("Enter a word: ")
	fmt.Scanf("%s ", &word)

	firstIdx := 0
	lastIdx := len(word) - 1

	isPalindrome := true
	for firstIdx < lastIdx {
		if word[firstIdx] != word[lastIdx] {
			isPalindrome = false
			break
		}

		firstIdx++
		lastIdx--
	}

	if isPalindrome {
		fmt.Printf("'%s' is a palindrome.", word)
	} else {
		fmt.Printf("'%s' is not a palindrome.", word)
	}

}
