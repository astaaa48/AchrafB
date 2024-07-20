package main

import (
	"fmt"
)


func DigitLen(n, base int) int {
count := 1
if n < 0 {
	n = n * -1
}
if base < 0 {
	return -1
}
for n/base > 0 {
count++
n = n/base
}

return count
}

func main() {
	fmt.Println(DigitLen(100, 10))
	fmt.Println(DigitLen(100, 2))
	fmt.Println(DigitLen(-100, 16))
	fmt.Println(DigitLen(100, -1))
}