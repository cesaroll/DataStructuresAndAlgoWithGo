package main

import "fmt"

func main() {
	nums := []int{3, -1, 4, 12, -8, 5, 6}
	fmt.Println(MaxSum(nums, 4))
}

func MaxSum(nums []int, k int) int {
	if k > len(nums) {
		return 0
	}

	sum := 0
	right := 0
	result := 0

	for ; right < k; right++ {
		result += nums[right]
	}

	for ; right < len(nums); right ++ {
		sum += nums[right] - nums[right - k]
		result = max(result, sum)
	}

	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}