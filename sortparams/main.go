package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	s := []string(os.Args[1:])

	var temp string
	len := len(s) - 1
	for i := 0; i < len; i++ {
		for j := 0; j < len-i; j++ {
			if s[j] > s[j+1] {
				temp = s[j]
				s[j] = s[j+1]
				s[j+1] = temp
			}
		}
	}
	for _, word := range s {
		for _, w := range word {
			z01.PrintRune(rune(w))
		}
		z01.PrintRune(10)
	}
}
