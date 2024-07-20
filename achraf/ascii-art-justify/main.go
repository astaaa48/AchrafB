package main
import (
	"fmt"
	"os"
	"strings"
  
	ascii "ascii/func"
)
// main function
func main() {
	args := os.Args
	len_args := len(args)
	flag := ""
	option := ""
	if len_args == 1 {
		return
	}
	if len_args > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		return
	}
	if len(args[1]) > 7 && args[1][:8] == "--align=" {
		flag, option = ascii.Get_flag(args[1])
        if option != "center" && option != "left"  && option != "right" && option != "justify" {
            fmt.Println("option can be onlt one of the type : -center-left-right-justify")
            return
        } 
		args = append(args[:1], args[2:]...)
		len_args--
	}
	if len(args[1]) > 9 && args[1][:9] == "--output=" {
		flag, option = ascii.Get_flag(args[1])
		args = append(args[:1], args[2:]...)
		len_args--
	}

	Word := ""
	File := "standard.txt"
	if len_args >= 2 {
		Word = args[1]
	}
	if len_args == 3 {
		File = args[2]
		if len(File) > 4 && File[len(File)-4:] != ".txt" {
			File = File + ".txt"
		}
	}
	Words := strings.Split(Word, "\\n")
	// reading File
	AsciiArtFile := ascii.ReadFiles("Assci/" + File)
	ascii.Checker(Words, AsciiArtFile, option, flag)


   
}