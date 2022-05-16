package piscine

func BasicAtoi(s string) int {
	n := 0
	for i, word := range s {
		if i != 0 {
			n = n * 10
		}
		n = n + int((word - '0'))
	}
	return n
}
