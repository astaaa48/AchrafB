package ascii

import (
	"bufio"
	"os"
)

func ReadFiles(namefile string) []string {
	file, err := os.Open(namefile)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	read := bufio.NewScanner(file)
	var arr []string
	i := 0
	for read.Scan() {
		arr = append(arr, read.Text())
		i++
	}
	return arr
}
