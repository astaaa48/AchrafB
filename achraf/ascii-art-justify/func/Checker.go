package ascii

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)


func Checker(words, ascii []string, option, flag string) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Failed to execute command:", err)
		return
	}

	size := strings.TrimSpace(string(out))
	parts := strings.Split(size, " ")
	width, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println("Failed to execute command:", err)
		return
	}

	spaceCount := 0
	allEmpty := true

	fmt.Println(words)
	for _, word := range words {
		for _, char := range word {
			if char == ' ' {
				spaceCount++
			}
		}

		// Check if any word is not empty
		if word != "" {
			allEmpty = false
		}
	}

	if allEmpty {
		fmt.Println("All words are empty.")
		return
	}

	res := ""
	for i := 0; i < len(words); i++ {
		if flag == "" || flag == "--align=" || flag == "--output=" {
			res += Printer(words[i], ascii)
		}

		if (words[i] == "" && i > 0) || (words[i] == "" && i == 0 && len(words) > 1 && !allEmpty) {
			if flag == "" || flag == "--align=" || flag == "--output=" {
				res += "\n"
			}
		}
	}

	if flag == "" {
		fmt.Printf("%s", res)
	} else if flag == "--align=" {
		printAligned(res, width, option)
	} else if flag == "--output=" {
		file, err := os.Create(option)
		if err != nil {
			fmt.Println("Failed to create file:", err)
			return
		}
		defer file.Close()
		file.WriteString(res)
	}
}

func printAligned(res string, width int, option string) {
	lines := strings.Split(res, "\n")
	var result string

	switch option {
	case "center":
		for _, line := range lines {
			if len(line) > width {
				fmt.Println("len of the word > terminal width")
				return
			}
			padding := (width - len(line)) / 2
			result += strings.Repeat(" ", padding) + line + "\n"
		}
	case "left":
		result = res
	case "right":
		for _, line := range lines {
			if len(line) > width {
				fmt.Println("len of the word > terminal width")
				return
			}
			padding := width - len(line)
			result += strings.Repeat(" ", padding) + line + "\n"
		}
	case "justify":
		for _, line := range lines {
			words := strings.Fields(line)
			if len(words) == 1 {
				result += line + "\n"
				continue
			}
			totalSpaces := width - len(strings.ReplaceAll(line, " ", ""))
			spacesBetweenWords := totalSpaces / (len(words) - 1)
			extraSpaces := totalSpaces % (len(words) - 1)
			for i, word := range words {
				result += word
				if i < len(words)-1 {
					result += strings.Repeat(" ", spacesBetweenWords)
					if i < extraSpaces {
						result += " "
					}
				}
			}
			result += "\n"
		}
	default:
		result = res
	}

	fmt.Print(result)
}
