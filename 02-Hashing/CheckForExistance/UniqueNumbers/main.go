package main

import "fmt"

func main() {

  fmt.Println(uniqueNumbers([]int{5,4,6,8}))

}

func uniqueNumbers(nums []int) []int {
  result := []int{}
  m := make(map[int]int)

  for i := 0; i < len(nums); i++ {
    x := nums[i]
    m[x] = i
  }

  for i := 0; i < len(nums); i++ {
    x := nums[i]
    _, exists := m[x+1]
    if !exists {
      _, exists := m[x-1]
      if !exists {
        result = append(result, x)
      }
    }
  }

  return result
}
