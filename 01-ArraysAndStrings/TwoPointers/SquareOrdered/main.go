package main

import "fmt"

func main() {
	// nums := []int{-7, -3, 2, 3, 11}
	// nums := []int{1}
	nums := []int{0, 2}

	fmt.Println(nums)
	fmt.Println(sortedSquares(nums))
}

func sortedSquares(nums []int) []int {
	n := len(nums)
	squares := make([]int, n)
	left, right := 0, n - 1

	for idx := n-1; idx >= 0; idx-- {
		square := 0

		if abs(nums[left]) < abs(nums[right]) {
			square = nums[right] * nums[right]
			right--
		} else {
			square = nums[left] * nums[left]
			left++
		}
		squares[idx] = square
	}

	return squares
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}