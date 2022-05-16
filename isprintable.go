package piscine

func IsPrintable(s string) bool {
	if len(s) == 0 {
		return false
	}
	a := []rune(s)
	for i := 0; i < len(a); i++ {
		if a[i] >= 32 && a[i] <= 127 {
		} else {
			return false
		}
	}
	return true
}
