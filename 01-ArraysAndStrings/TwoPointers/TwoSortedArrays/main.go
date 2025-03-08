package main

import (
	"fmt"
)

func main() {

	array1 := []int{1, 2, 4, 6, 8, 9, 14, 15, 20, 25}
	array2 := []int{2, 5, 6, 7, 8, 12, 14, 18}

	result := merge(array1, array2)
	fmt.Println(result)
}

func merge(array1 []int, array2 []int) []int {
	idx1 := 0
	idx2 := 0
	result := []int{}

	for idx1 < len(array1) && idx2 < len(array2) {
		value1 := array1[idx1]
		value2 := array2[idx2]

		if value1 <= value2 {
			result = append(result, value1)
			idx1++
		} else {
			result = append(result, value2)
			idx2++
		}
	}

	if idx1 < len(array1) {
		result = append(result, array1[idx1:]...)
	} else if idx2 < len(array2) {
		result = append(result, array2[idx2:]...)
	}

	return result
}
