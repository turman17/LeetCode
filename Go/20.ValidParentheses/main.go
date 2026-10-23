package main

import (
	"fmt"
)

func isOpened(s rune) bool {
	return s == '(' || s == '{' || s == '['
}

func isMatching(open, close rune ) bool{
	switch open{
		case '(':
			return close == ')'
		case '{':
			return close == '}'	
		case '[':
			return close == ']'
	}
	return false
}

func isValid(s string) bool {
	if len(s) % 2 != 0 {
		return false
	}
	stack := []rune{}
	for _ , char := range(s) {
		if isOpened(char){
			stack = append(stack, char)	
		} else {
			if len(stack) != 0 && isMatching (rune(stack[len(stack) - 1]), char) {
				stack = stack[:len(stack)-1]
			}else {
				return false
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

func main(){
	s := "))"
	fmt.Println(isValid(s))
}
