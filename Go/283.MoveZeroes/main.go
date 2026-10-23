package main

import "fmt"

func moveZeroes(nums []int) {
	j := 0
	for i := range nums {
		j = i
		if nums[i] == 0 {
			for j < len(nums) && nums[j] == 0 {
				j++
			}
			if j < len(nums) && nums[j] != 0 {
				nums[i] = nums[j]
				nums[j] = 0
			}
		}
	}
}

func main() {
	nums := []int{1, 0, 2, 3, 0, 4}
	moveZeroes(nums)

	fmt.Print(nums)
}
