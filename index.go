package piscine

func Index(s string, toFind string) int {
	match := false

	a := []rune(s)
	b := []rune(toFind)
	if len(a) == 0 {
		return -1
	}
	if len(b) == 0 {
		return 0
	}
	if len(b) > len(a) {
		return -1
	}
	for i := 0; i < len(a)-len(b); i++ {
		if a[i] == b[0] {
			match = true
			for j := 0; j < len(b)-1; j++ {
				if a[i+j] != b[j] {
					match = false
				}
			}
			if match {
				return i
			}
		}
	}
	return -1
}
