package main

import "fmt"

func moveZeroes(nums []int) {
	j := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}

func main() {
	nums := []int{1, 0, 2, 3, 0, 4}
	moveZeroes(nums)

	fmt.Print(nums)
}
