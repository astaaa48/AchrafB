package ascii
import "fmt"
func AsciiCheker(word []byte) bool {
	for j := 0; j < len(word); j++ {
		if word[j] < 32 || word[j] > 126 {
			fmt.Println("You can only use ASCII characters.")
			return true
		}
	}
	return false
}