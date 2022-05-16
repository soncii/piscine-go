package piscine

func NRune(s string, n int) rune {
	ss := []rune(s)
	if n > len(s) {
		return 0
	}
	if n <= 0 {
		return 0
	}
	return ss[n-1]
}
