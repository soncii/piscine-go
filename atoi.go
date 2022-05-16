package piscine

func Atoi(s string) int {
	if s == "" {
		return 0
	}
	for i, word := range s {
		if word > '9' || word < '0' {
			if i == 0 && (word == '-' || word == '+') {
			} else {
				return -123456
			}
		}
	}
	if s[0] == '-' || s[0] == '+' {
		n := 0
		for i, word := range s {
			if i != 0 {
				if i != 1 {
					n = n * 10
				}
				n = n + int((word - '0'))
			}
		}
		if s[0] == '-' {
			return n * (-1)
		}
		return n
	}
	n := 0
	for i, word := range s {
		if i != 0 {
			n = n * 10
		}
		n = n + int((word - '0'))
	}
	return n
}
