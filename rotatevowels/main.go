package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var global string
	if len(os.Args) < 2 {
		z01.PrintRune('\n')
	} else {
		for i := 1; i < len(os.Args); i++ {
			global += os.Args[i]

			global += " "

		}
		ind := []int{}
		g := []rune(global)
		for i := 0; i < len(g); i++ {
			if g[i] == 'A' || g[i] == 'E' || g[i] == 'O' || g[i] == 'U' || g[i] == 'I' || g[i] == 'a' || g[i] == 'e' || g[i] == 'o' || g[i] == 'u' || g[i] == 'i' {
				ind = append(ind, i)
			}
		}
		for i := 0; i < len(ind)/2; i++ {
			g[ind[i]], g[ind[len(ind)-i-1]] = g[ind[len(ind)-i-1]], g[ind[i]]
		}
		for i := 0; i < len(g); i++ {
			z01.PrintRune(g[i])
		}
		z01.PrintRune('\n')
	}
}
