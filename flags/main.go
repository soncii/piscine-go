package main

import (
	"fmt"
	"os"
)

func main() {
	s := []string(os.Args)
	strappend := []rune{}

	if len(s) == 1 {
		fmt.Println("--insert")
		fmt.Println("  -i")
		fmt.Println("	 This flag inserts the string into the string passed as argument.")
		fmt.Println("--order")
		fmt.Println("  -o")
		fmt.Println("	 This flag will behave like a boolean, if it is called it will order the argument.")

	} else if s[1] == "--help" || s[1] == "-h" {
		fmt.Println("--insert")
		fmt.Println("  -i")
		fmt.Println("	 This flag inserts the string into the string passed as argument.")
		fmt.Println("--order")
		fmt.Println("  -o")
		fmt.Println("	 This flag will behave like a boolean, if it is called it will order the argument.")

	} else {
		first := []rune(s[1])
		ins := []rune("--insert=")
		mainstr := []rune{}
		insert := true
		order := false
		for i := 0; i <= 8; i++ {
			if ins[i] != first[i] {
				insert = false
				break
			}
		}
		index := 9
		if first[0] == '-' && first[1] == 'i' && first[2] == '=' {
			insert = true
			index = 3
		}
		if insert {
			for i := index; i < len(first); i++ {
				strappend = append(strappend, first[i])
			}
			if s[2] == "-o" || s[2] == "--order" {
				order = true
				mainstr = []rune(s[3])
			} else {
				mainstr = []rune(s[2])
			}
		} else {
			if s[1] == "-o" || s[1] == "--order" {
				order = true
				mainstr = []rune(s[2])
			} else {
				mainstr = []rune(s[1])
			}
		}
		mainstr = append(mainstr, strappend...)
		if order {
			len := len(mainstr) - 1
			for i := 0; i < len; i++ {
				for j := 0; j < len-i; j++ {
					if mainstr[j] > mainstr[j+1] {
						temp := mainstr[j]
						mainstr[j] = mainstr[j+1]
						mainstr[j+1] = temp
					}
				}
			}
		}
		fmt.Println(string(mainstr))
	}
}
