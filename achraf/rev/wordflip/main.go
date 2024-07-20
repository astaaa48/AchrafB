package main

import (
	"fmt"
	"strings"
)

func WordFlip(str string) string {
	str = strings.TrimSpace(str)
	if str == "" {
		return "Invalid Output\n"
	}

	words := strings.Fields(str)
	rs := ""

	for i := len(words) - 1; i >= 0; i-- {
		if i != len(words) - 1 {
			rs += " "
		}
		rs += words[i]
	}
	
	return rs + "\n"
}

func main() {
	fmt.Print(WordFlip("First second last"))
	fmt.Print(WordFlip(""))
	fmt.Print(WordFlip("     "))
	fmt.Print(WordFlip(" hello  all  of  you! "))
}
