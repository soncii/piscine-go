package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "--upper" {
			s := os.Args[2:]
			for i := 0; i < len(s); i++ {
				n := 0
				for _, w := range s[i] {
					if w >= 48 && w <= 57 {
						n = n * 10
						n = n + int(w-'0')
					} else {
						z01.PrintRune(' ')
						break
					}
				}
				if n > 0 && n <= 26 {
					z01.PrintRune(rune(n + 64))
				} else {
					z01.PrintRune(' ')
				}
			}
		} else {
			s := os.Args[1:]
			for i := 0; i < len(s); i++ {
				n := 0
				for _, w := range s[i] {
					if w >= 48 && w <= 57 {
						n = n * 10
						n = n + int(w-'0')
					} else {
						n = -1
						break
					}
				}
				if n > 0 && n <= 26 {
					z01.PrintRune(rune(n + 96))
				} else {
					z01.PrintRune(' ')
				}
			}

		}
		z01.PrintRune(10)
	}
}
