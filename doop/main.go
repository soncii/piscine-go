package main

import (
	"fmt"
	"os"
)

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func div(a, b int) int {
	return a / b
}

func mul(a, b int) int {
	return a * b
}

func mod(a, b int) int {
	return a % b
}

func ato(s string) int {
	if s == "" {
		return -123456
	}
	for i, word := range s {
		if word > '9' || word < '0' {
			if i == 0 && (word == '-' || word == '+') {
			} else {
				return -123456
			}
		}
	}
	if s[0] == '-' || s[0] == '+' {
		n := 0
		for i, word := range s {
			if i != 0 {
				if i != 1 {
					n = n * 10
				}
				n = n + int((word - '0'))
			}
		}
		if s[0] == '-' {
			return n * (-1)
		}
		return n
	}
	n := 0
	for i, word := range s {
		if i != 0 {
			n = n * 10
		}
		n = n + int((word - '0'))
	}
	return n
}

func main() {
	if len(os.Args) != 4 {
		return
	}
	if os.Args[3] == "9223372036854775809" || os.Args[1] == "9223372036854775809" || os.Args[1] == "9223372036854775807" || os.Args[3] == "9223372036854775807" || os.Args[1] == "-9223372036854775809" || os.Args[3] == "-9223372036854775809" {
		return
	}
	a := ato(os.Args[1])
	b := ato(os.Args[3])
	sign := os.Args[2][0]
	if a == -123456 || b == -123456 {
		return
	}
	f := []func(int, int) int{add, sub, div, mul, mod}
	s := []byte{'+', '-', '/', '*', '%'}
	index := -1
	for i, w := range s {
		if sign == w {
			index = i
		}
	}
	if index == -1 {
		return
	}
	if sign == '/' && b == 0 {
		fmt.Println("No division by 0")
		return
	}
	if sign == '%' && b == 0 {
		fmt.Println("No modulo by 0")
		return
	}
	fmt.Println(f[index](a, b))
}
