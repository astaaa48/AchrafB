package main

import (
	"fmt"
	
)

func Itoa(n int) string {
rs:=""
signe := 1
if n < 0 {
	signe = -1
	n = n * signe
	
}
if n == 0{
	return "0"
}

for n!= 0{
	
	rs = string(n%10 + 48) + rs 
	n = n/10
}



if signe < 0 {
	rs = "-" + rs 
}

return rs



}

func main() {
    fmt.Println(Itoa(12345))
    fmt.Println(Itoa(0))
    fmt.Println(Itoa(-1234))
    fmt.Println(Itoa(987654321))
}