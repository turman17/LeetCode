package main

import "fmt"

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currentSum := maxSum
	for i := 1; i < len(nums); i++ {
		if (currentSum + nums[i]) >= nums[i] {
			currentSum += nums[i]
		} else {
			currentSum = nums[i]
		}
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}

func main() {
	nums := []int{-1,-2}
	fmt.Println(maxSubArray(nums))
}
