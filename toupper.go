package piscine

func ToUpper(s string) string {
	a := []rune(s)
	for i := 0; i < len(a); i++ {
		if a[i] > 96 && a[i] < 123 {
			a[i] = a[i] - 32
		}
	}
	return string(a)
}
