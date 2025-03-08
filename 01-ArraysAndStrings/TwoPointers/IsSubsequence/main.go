package main

import "fmt"

func main() {
	fmt.Println(isSubsequence([]rune("ace"), []rune("abcde")))
	fmt.Println(isSubsequence([]rune("aec"), []rune("abcde")))
}

func isSubsequence(slice1 []rune, slice2 []rune) bool {

	idx1 := 0
	idx2 := 0

	for idx1 < len(slice1) && idx2 < len(slice2) {
		val1 := slice1[idx1]
		val2 := slice2[idx2]

		if val1 == val2 {
			if idx1 == len(slice1)-1 {
				return true
			}
			idx1++
			idx2++
		} else {
			idx2++
		}
	}

	return false
}
