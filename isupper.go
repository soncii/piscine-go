package piscine

func IsUpper(s string) bool {
	if len(s) == 0 {
		return false
	}
	a := []rune(s)
	for i := 0; i < len(a); i++ {
		if a[i] > 64 && a[i] < 91 {
		} else {
			return false
		}
	}
	return true
}
