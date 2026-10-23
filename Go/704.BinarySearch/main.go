package main

import (
	"fmt"
)

func calcMiddlePos(left, right int) int {
	return (right + left) / 2
}

func search(nums []int, target int) int {
	lastPos := len(nums) - 1
	firstPos := 0
	for firstPos <= lastPos {
		middlePos := calcMiddlePos(firstPos, lastPos)
		if nums[middlePos] == target {
			return middlePos
		} else if nums[middlePos] > target {
			lastPos = middlePos-1
		}else {
			firstPos = middlePos+1
		}
	}
	return -1
}

func main() {
	nums := []int{-1, 0, 3, 4, 5, 9}
	target := 3

	fmt.Println(search(nums, target))
}
