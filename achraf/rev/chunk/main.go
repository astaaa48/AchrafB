package main

import (
	"fmt"
)

func Chunk(slice []int, size int) {
	if size == 0 {
		fmt.Println("\n")
		return
	}

var rs [][]int
	

for size < len(slice) {
	rs = append(rs, slice[size:])
}

fmt.Println(rs)
	
}

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}



