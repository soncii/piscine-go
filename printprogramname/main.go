package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	s := os.Args[0]
	for i, word := range s {
		if i != 0 && i != 1 {
			z01.PrintRune(rune(word))
		}
	}
	z01.PrintRune(10)
}
