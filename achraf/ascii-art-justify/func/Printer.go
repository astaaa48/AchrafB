package ascii
func Printer(word string, ascii []string) string {
	data := ""
	if len(word) == 0 {
		// protecting the case of empty element that means there new-line
		return ""
	}
	byt := []byte(word)
	// Check if all words are ASCII.
	if AsciiCheker(byt) {
		return ""
	}
	// the looping for all the 8 lines
	for i := 0; i < 8; i++ {
		// this loop print the single line
		
		for j := 0; j < len(byt); j++ {
			data += ascii[(int(byt[j])-32)*9+1+i]
		}
		data += "\n"
	}
	return data
}