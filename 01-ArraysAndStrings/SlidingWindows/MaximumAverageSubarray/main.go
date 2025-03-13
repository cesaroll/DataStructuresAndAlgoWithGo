package main

import "fmt"

func main() {

	fmt.Println(findMaxAverage([]int{1,12,-5,-6,50,3}, 4))

}

func findMaxAverage(nums []int, k int) float64 {
    left := 0
    right := 0
    sum := 0
    var avg float64 = 0
    
    for ; right < k; right++ {
        sum += nums[right]
    }
    
    avg = float64(sum) / float64(k)
    
    for ; right < len(nums); right++ {
        sum += nums[right]
        
        if right - left + 1 > k {
            sum -= nums[left]
            left++
        }
        
        if float64(sum) / float64(k) > avg {
            avg = float64(sum) / float64(k)
        }
    }
    
    return avg
}