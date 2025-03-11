package main

import "fmt"

func main() {
	fmt.Println(longestSubArrayOneFlip("1101100111"))
}

// Returns the size of the longest subarray with max one flip
func longestSubArrayOneFlip(s string) int {

	left := 0
	zeroCount := 0
	result := 0

	for right := 0; right < len(s); right++ {
		if s[right] == '0' {
			zeroCount++
		}

		for zeroCount > 1 {
			if s[left] == '0' {
				zeroCount--
			}
			left++
		}

		result = max(result, right-left+1)
	}
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
