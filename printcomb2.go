package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 48; i <= 57; i++ {
		for j := 48; j <= 57; j++ {
			for k := i; k <= 57; k++ {
				if k == i && j != 57 {
					for g := j + 1; g <= 57; g++ {
						z01.PrintRune(rune(i))
						z01.PrintRune(rune(j))
						z01.PrintRune(rune(' '))
						z01.PrintRune(rune(k))
						z01.PrintRune(rune(g))
						if i == 57 && j == 56 && k == 57 && g == 57 {
							z01.PrintRune(rune('\n'))
						} else {
							z01.PrintRune(rune(','))
							z01.PrintRune(rune(' '))
						}

					}
				}
				if k == i && j == 57 && k != 57 {
					k++
				}
				if i == 57 && j == 56 {
				} else {
					if k != i {
						for g := 48; g <= 57; g++ {
							z01.PrintRune(rune(i))
							z01.PrintRune(rune(j))
							z01.PrintRune(rune(' '))
							z01.PrintRune(rune(k))
							z01.PrintRune(rune(g))
							if i == 57 && j == 56 && k == 57 && g == 57 {
							} else {
								z01.PrintRune(rune(','))
								z01.PrintRune(rune(' '))
							}

						}
					}
				}
			}
		}
	}
}
