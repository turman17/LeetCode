package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
    raw := len(nums)
    new := make(map[int]int,raw)
    for i, value := range nums{
    	new[value] = i
    }
    if raw == len(new) {
    	return false
    }
    return true
}

func main(){
	nums := []int{1,2,3,4}
	fmt.Println(containsDuplicate(nums))
}