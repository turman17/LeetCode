package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	j := len(s) - 1
	for i := 0; i < j; {
		if unicode.IsLetter(rune(s[i])) || unicode.IsDigit(rune(s[i])) &&
			unicode.IsLetter(rune(s[j])) || unicode.IsDigit(rune(s[j])) {
			if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) ||
				unicode.IsDigit(rune(s[i])) != unicode.IsDigit(rune(s[j])) {
				return false
			}
			i++
			j--
		} else {
			if !unicode.IsLetter(rune(s[i])) &&
				!unicode.IsDigit(rune(s[i])) {
				i++
			}
			if !unicode.IsLetter(rune(s[j])) &&
				!unicode.IsDigit(rune(s[j])) {
				j--
			}
		}
	}
	return true
}

func main() {
	s := "0p"

	fmt.Println(isPalindrome(s))
}
