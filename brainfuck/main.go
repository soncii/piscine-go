package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	s := os.Args[1]
	out := [400]int{}
	loop := 0
	i := 0
	for j := 0; j < len(s); j++ {
		w := s[j]
		if w == '+' {
			out[i] += 1
		}
		if w == '>' {
			i++
		}
		if w == '<' {
			i--
		}
		if w == '.' {
			z01.PrintRune(rune(out[i]))
		}
		if w == '[' {
			if out[i] == 0 {
				for s[j] != ']' && j < len(s) {
					j++
				}
				j++
			} else {
				loop = j
			}
		}
		if w == ']' {
			if out[i] != 0 {
				j = loop
			}
		}
		if w == '-' {
			out[i]--
		}
	}
}
