package piscine

func SplitWhiteSpaces(s string) []string {
	n := []rune(s)
	temp := []rune{}
	res := []string{}
	for i := 0; i < len(n); i++ {
		if n[i] != rune(32) && n[i] != rune(9) && n[i] != '\n' {
			temp = append(temp, n[i])
		} else {
			if string(temp) != "" {
				res = append(res, string(temp))
			}
			temp = []rune("")
		}
	}
	res = append(res, string(temp))
	return res
}
