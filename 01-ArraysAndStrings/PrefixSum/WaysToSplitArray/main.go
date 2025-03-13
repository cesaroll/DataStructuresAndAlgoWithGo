package main

import "fmt"

func main () {
	fmt.Println(waysToSplitArray([]int{10,4,-8,7}))
}

func waysToSplitArray(nums []int) int {
	leftSum := nums[0]
	rightSum := 0
	result := 0

	for i:=1; i<len(nums); i++ {
		rightSum += nums[i]
	}
	fmt.Println(rightSum)

	// [10,  4, -8,  7]
    //  --  __________ 10, 3
    // [10, 14,  6, 13]

	for i:=0; i<len(nums)-1; i++ {
		if leftSum >= rightSum {
			result++
		}
		fmt.Printf("%d - %d\n", leftSum, rightSum)
		leftSum += nums[i+1]
		rightSum -= nums[i+1]
	}

	return result
}