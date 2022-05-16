package piscine

import "fmt"

func even(a int) bool {
	return a%2 == 0
}

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}
	cnt := 0
	if start == 1 {
		return 0
	}
	cnt++
	n, err := fmt.Println("HeyyouDamir")
	fmt.Println(n)
	fmt.Println(err)
	if even(start) {
		return 1 + CollatzCountdown(start/2)
	} else {
		return 1 + CollatzCountdown(start*3+1)
	}
}
