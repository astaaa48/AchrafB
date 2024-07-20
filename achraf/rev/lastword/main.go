package main

import (
	"fmt"
	"strings"
	
)

func LastWord(s string) string {
	
	s = strings.TrimSpace(s)
	
	words := strings.Fields(s)
	
	if len(words) == 0 {
		return "" + "\n"
	}

	return words[len(words)-1] + "\n"
}

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))

	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
	
}