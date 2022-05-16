package main

import (
	"fmt"
	"os"
)

func atoi(s string) int {
	if s == "" {
		return 0
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
	er := false
	arg := os.Args[3:]
	for i, name := range arg {
		file, err := os.Open(name)
		if err != nil {
			out := "open " + name + ": no such file or directory\n"
			fmt.Printf("%s", out)
			er = true
			continue
		}
		c := atoi(os.Args[2])
		arr := make([]byte, 1000)
		n, _ := file.Read(arr)
		if i != 0 {
			fmt.Printf("\n")
		}
		if (n - c) > 0 {
			fmt.Printf("==> %s <==\n", name)
			fmt.Printf("%s", string(arr[n-c:n]))
		} else {
			fmt.Printf("==> %s <==\n", name)
			fmt.Printf("%s", string(arr[:n]))
			os.Exit(1)
		}
	}
	if er {
		os.Exit(1)
	}
}
