package main

import (
	"fmt"
)

func CheckNumber(arg string) bool {
count := 0
for _, char := range arg{
	if char > '0' && char < '9'{
		count++
	} 
}
if count > 0{
	return true
}else{
	return false
}

}

func main() {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
}