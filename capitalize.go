package piscine

func Capitalize(s string) string {
	a := []rune(s)
	word := false
	for i := 0; i < len(s); i++ {
		if (a[i] > 64 && a[i] < 91) && !word {
			word = true
		} else if (a[i] > 64 && a[i] < 91) && word {
			a[i] = a[i] + 32
		} else if (a[i] > 96 && a[i] < 123) && word {
		} else if (a[i] > 96 && a[i] < 123) && !word {
			a[i] = a[i] - 32
			word = true
		} else if a[i] >= 48 && a[i] <= 57 {
			word = true
		} else {
			word = false
		}
	}
	return string(a)
}
