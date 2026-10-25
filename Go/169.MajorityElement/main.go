package main

import "fmt"

func majorityElement(nums []int) int {
	count := make(map[int]int)
	for _ , num := range nums {
		count[num]++
		if count[num] > (len(nums)/2){
			return num
		}
	}
	return 0
}

func main() {
	nums := []int{2, 2,1 , 1, 1, 1, 1, 2, 2}
	fmt.Print(majorityElement(nums))
}
