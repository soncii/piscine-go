package piscine

func IsNumeric(s string) bool {
	if len(s) == 0 {
		return false
	}
	a := []rune(s)
	for i := 0; i < len(a); i++ {
		if a[i] >= 48 && a[i] <= 57 {
		} else {
			return false
		}
	}
	return true
}
