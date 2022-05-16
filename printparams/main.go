package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	s := os.Args[1:]
	for _, w := range s {
		for _, word := range w {
			z01.PrintRune(rune(word))
		}
		z01.PrintRune(10)
	}
}
