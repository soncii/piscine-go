package piscine

func IsLower(s string) bool {
	a := []rune(s)
	if len(s) == 0 {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] > 96 && a[i] < 123 {
		} else {
			return false
		}
	}
	return true
}
