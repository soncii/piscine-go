package piscine

func StrRev(s string) string {
	final := []rune(s)
	len := 0
	for range s {
		len++
	}
	len--
	for i, word := range s {
		final[len-i] = word
	}
	r := string(final)
	return r
}
