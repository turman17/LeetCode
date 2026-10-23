package main

import "fmt"

func moveZeroes(nums []int) {
	j := 0
	for i := range nums {
		if nums[j] == 0 && nums[i] != 0 {
			nums[j] = nums[i]
			nums[i] = 0
			j++
		} else if nums[j] != 0 {
			j++
		}
	}
}

func main() {
	nums := []int{1, 0, 2, 3, 0, 4}
	moveZeroes(nums)

	fmt.Print(nums)
}
