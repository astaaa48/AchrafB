package main

import (
	"fmt"
	"strconv"
)

func ZipString(s string) string {
	if len(s) == 0 {
		return ""
	}

	var result string
	count := 1

	for i := 1; i < len(s); i++ {
		// Only process Latin alphabet characters
		if (s[i] < 'a' || s[i] > 'z') && (s[i] < 'A' || s[i] > 'Z') {
			result += strconv.Itoa(count) + string(s[i-1])
			result += "1" + string(s[i])
			count = 1
		} else if s[i] == s[i-1] {
			count++
		} else {
			result += strconv.Itoa(count) + string(s[i-1])
			count = 1
		}
	}

	// Add the last group of characters
	result += strconv.Itoa(count) + string(s[len(s)-1])

	return result
}

func main() {
	fmt.Println(ZipString("22"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog")) // 1T1h2e1 1q2u1i1c1k1 1b1r1o2w1n1 1f1o1x1 1j2u1m1p1s1 1o1v1e1r1 1t1h1e1 1l3a1z1y1 1d1o1g
	fmt.Println(ZipString("Helloo Therre!"))                      // 1H1e2l2o1 1T1h1e2r1e1!
}