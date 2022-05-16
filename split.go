package piscine

func ind(s string, tofind string, index *int) string {
	match := false

	a := []rune(s)
	b := []rune(tofind)
	n := []rune{}

	for i := *index; i < len(a)-len(b); i++ {
		if a[i] == b[0] {
			match = true
			for j := 0; j < len(b); j++ {
				if a[i+j] != b[j] {
					match = false
				}
			}
			if match {
				for j := *index; j < i; j++ {
					n = append(n, a[j])
				}
				*index = i + len(b)
				return string(n)
			}
		}
	}
	for i := *index; i < len(a); i++ {
		n = append(n, a[i])
	}
	*index = len(a)
	return string(n)
}

func Split(s, sep string) []string {
	index := 0

	a := []string{}

	for index < len(s) {
		a = append(a, ind(s, sep, &index))
	}
	return a
}
