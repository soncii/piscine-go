package piscine

func check(s []rune, b []rune) bool {
	l := len(b)
	if l < 2 {
		return true
	}
	for i := 0; i < l; i++ {
		if b[i] == '+' || b[i] == '-' {
			return true
		}
		for j := i + 1; j < l; j++ {
			if b[i] == b[j] {
				return true
			}
		}
	}
	for _, word := range s {
		in := false
		for _, w := range b {
			if word == w {
				in = true
			}
		}
		if !in {
			return true
		}
	}
	return false
}

func AtoiBase(s string, base string) int {
	b := []rune(base)
	num := []rune(s)
	if check(num, b) {
		return 0
	}
	l := len(b)
	//	bas := len(b)
	n := 0
	pow := 0
	digit := 0
	for i := len(num) - 1; i >= 0; i-- {
		for j := 0; j < len(b); j++ {
			if num[i] == b[j] {
				digit = j
			}
		}
		n = n + digit*IterativePower(l, pow)
		pow++
	}
	return n
}
