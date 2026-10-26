package main

import "fmt"

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] == 10{
			digits[i] = 0
		} else{
			return digits
		}
	}
	return append([]int{1}, digits...)
}

func main() {
	digits := []int{1,9,9}
	fmt.Print(plusOne(digits))
}
