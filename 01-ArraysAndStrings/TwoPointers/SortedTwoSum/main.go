package main

func main() {
	nums := []int{1, 2, 4, 6, 8, 9, 14, 15}
	target := 13
	result := twoSumOrdered(nums, target)
	if result {
		println("Two sum exists")
	} else {
		println("Two sum does not exist")
	}
}

func twoSumOrdered(nums []int, target int) bool {

	leftIdx := 0
	rightIdx := len(nums) - 1

	for leftIdx < rightIdx {
		sum := nums[leftIdx] + nums[rightIdx]
		if sum == target {
			return true
		}
		if sum < target {
			leftIdx++
		} else {
			rightIdx--
		}
	}
	return false
}
