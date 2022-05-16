package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == -9223372036854775808 {
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

	} else {
		if n == 0 {
			z01.PrintRune(rune('0'))
		} else {
			var arr [32]int
			if n < 0 {
				z01.PrintRune(rune('-'))
				n = n * (-1)
			}

			var i int = 0
			for n > 10 {
				arr[i] = (n % 10) + 48
				i++
				n = n / 10
			}
			arr[i] = n + 48
			var j int = 0
			for j = i; j >= 0; j-- {
				z01.PrintRune(rune(arr[j]))
			}
		}
	}
}
