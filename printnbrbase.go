package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if len(base) < 2 {
		PrintStr("NV")
		return
	}
	if nbr == -9223372036854775808 {
		z01.PrintRune(rune('-'))
		z01.PrintRune(rune('9'))
		z01.PrintRune(rune('2'))
		z01.PrintRune(rune('2'))
		z01.PrintRune(rune('3'))
		z01.PrintRune(rune('3'))
		z01.PrintRune(rune('7'))
		z01.PrintRune(rune('2'))
		z01.PrintRune(rune('0'))
		z01.PrintRune(rune('3'))
		z01.PrintRune(rune('6'))
		z01.PrintRune(rune('8'))
		z01.PrintRune(rune('5'))
		z01.PrintRune(rune('4'))
		z01.PrintRune(rune('7'))
		z01.PrintRune(rune('7'))
		z01.PrintRune(rune('5'))
		z01.PrintRune(rune('8'))
		z01.PrintRune(rune('0'))
		z01.PrintRune(rune('8'))
		return
	}
	ind := 0
	b := []rune(base)
	l := len(b)
	minus := false
	if nbr < 0 {
		minus = true
		nbr = nbr * -1
	}

	for i := 0; i < l; i++ {
		if b[i] == '+' || b[i] == '-' {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
		for j := i + 1; j < l; j++ {
			if b[i] == b[j] {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}

	s := []rune{}
	d := 0
	for nbr >= l {
		d++
		if ind == 1 {
			s = append(s, b[nbr%l+1])
		}
		s = append(s, b[nbr%l])
		nbr = nbr / l
	}
	if ind == 1 {
		s = append(s, b[nbr+1])
	} else {
		s = append(s, b[nbr])
	}
	if minus {
		z01.PrintRune('-')
	}
	for i := d; i >= 0; i-- {
		z01.PrintRune(s[i])
	}
}
