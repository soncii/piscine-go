package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	i := 0
	if n == 0 {
		z01.PrintRune('0')
	} else {
		var a []int
		for n > 0 {
			a = append(a, n%10)
			n = n / 10
			i++
		}
		SortIntegerTable(a[:])
		for j := 0; j < i; j++ {
			z01.PrintRune(rune(a[j]) + 48)
		}
	}
}
