package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))
	currentVal := 1
	answer[0] = 1
	for i := 1; i < len(nums); i++ {
		currentVal *= nums[i-1]
		answer[i] = currentVal
	}
	currentVal = 1
	for i := len(nums) - 2; i >= 0; i-- {
		currentVal *= nums[i+1]
		answer[i] *= currentVal 
	}
	return answer
}

func main() {
	nums := []int{1, 2, 3, 4}

	fmt.Println(productExceptSelf(nums))
}
