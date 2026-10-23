package main

import "fmt"


func twoSum(nums []int, target int) []int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	var nums = []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum(nums, target))
	print()
}