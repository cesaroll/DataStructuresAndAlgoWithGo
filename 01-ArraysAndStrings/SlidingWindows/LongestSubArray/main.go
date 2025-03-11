package main

import "fmt"

func main() {
	// nums := []int{3, 2, 1, 3, 1, 1}
	nums := []int{6, 6, 6}
	fmt.Println(longestSubArray(nums, 4))
}

// Longest subarray whose sum is not bigger than k
func longestSubArray(nums []int, k int) int {

	left := 0
	sum := 0
	result := 0

	for right := 0; right < len(nums); right++ {
		sum += nums[right]

		for sum > k {
			sum -= nums[left]
			left++
		}

		result = max(result, right - left +1)

	}
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}