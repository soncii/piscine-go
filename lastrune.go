package piscine

func LastRune(s string) rune {
	ss := []rune(s)
	count := 0
	for range ss {
		count++
	}
	return ss[count]
}
