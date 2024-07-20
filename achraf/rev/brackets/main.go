package main

import (
	"fmt"
	"os"
)

func isBalanced(s string) bool {
	stack := []rune{}
	brackets := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stack = append(stack, char)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != brackets[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	if len(os.Args) == 1 {
		return
	}

	for _, arg := range os.Args[1:] {
		if isBalanced(arg) {
			fmt.Println("OK")
		} else {
			fmt.Println("Error")
		}
	}
}