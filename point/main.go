package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setpoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	var points point

	setpoint(&points)
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	dig1, dig2 := points.x/10, points.x%10
	z01.PrintRune(rune(dig1 + 48))
	z01.PrintRune(rune(dig2 + 48))
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	dig1, dig2 = points.y/10, points.y%10

	z01.PrintRune(rune(dig1 + 48))
	z01.PrintRune(rune(dig2 + 48))
	z01.PrintRune(rune('\n'))
}
