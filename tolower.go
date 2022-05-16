package piscine

func ToLower(s string) string {
	a := []rune(s)
	for i := 0; i < len(a); i++ {
		if a[i] > 64 && a[i] < 91 {
			a[i] = a[i] + 32
		}
	}
	return string(a)
}
