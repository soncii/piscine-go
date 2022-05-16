package piscine

func BasicAtoi2(s string) int {
	for _, word := range s {
		if word > '9' || word < '0' {
			return 0
		}
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
