package ascii
func Get_flag(s string) (option, flag string) {
	if len(s) > 7 && s[:8] == "--align=" {
		flag = "--align="
		option = s[8:]
	} else if len(s) > 8 && s[:9] == "--output=" {
		flag = "--output="
		option = s[9:]
	}
	return flag, option
}