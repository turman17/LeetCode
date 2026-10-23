package main

import (
	"fmt"
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}
	sSplit := strings.Split(s, "")
	tSplit := strings.Split(t, "")
	sort.Strings(sSplit)
	sort.Strings(tSplit)
	for i := range sSplit {
		if sSplit[i] != tSplit[i] {
			return false
		}
	}
	return true
}

func main() {
	s := "rat"
	t := "rats"

	fmt.Println(isAnagram(s, t))
}
