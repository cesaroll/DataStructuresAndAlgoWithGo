package main

import (
	// "fmt"
	"log"
)

func main() {

	var words = []string{"racecar", "cesar", "anna",
		"level", "avid diva", "a santa at nasa"}

	for i := 0; i < len(words); i++ {
		word := words[i]
		isPalindrome := isPalindrome(word)
		log.Printf("%s is Palindrome: %t\n", word, isPalindrome)
	}
}

func isPalindrome(word string) bool {

	i := 0
	j := len(word) - 1

	for i < j {
		left := word[i]
		right := word[j]
		if left != right {
			return false
		}
		i++
		j--
	}

	return true
}
