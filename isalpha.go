package piscine

func IsAlpha(s string) bool {
	if len(s) == 0 {
		return true
	}
	a := []rune(s)
	for i := 0; i < len(a); i++ {
		if (a[i] > 64 && a[i] < 91) || (a[i] > 96 && a[i] < 123) || (a[i] >= 48 && a[i] <= 57) {
		} else {
			return false
		}
	}
	return true
}
