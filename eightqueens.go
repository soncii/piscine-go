package piscine

import (
	"github.com/01-edu/z01"
)

func swap(s []int, i, j int) {
	temp := s[i]
	s[i] = s[j]
	s[j] = temp
}

func checking(s []int) bool {
	for i := 0; i < 8; i++ {
		// lower right diagonal
		ii := i + 1
		j := s[i]
		for ii+1 <= 8 && j+1 <= 8 {
			ii++
			j++
			if s[ii-1] == j {
				return false
			}
		}
		// upper right diagonal
		ii = i + 1
		j = s[i]
		for ii-1 > 8 && j+1 <= 8 {
			ii--
			j++
			if s[ii-1] == j {
				return false
			}
			// upper left diagonal
			ii = i + 1
			j = s[i]
		}
		for ii-1 > 0 && j-1 > 0 {
			ii--
			j--
			if s[ii-1] == j {
				return false
			}

		}
		// lower left diagonal
		ii = i + 1
		j = s[i]
		for ii+1 <= 8 && j-1 > 0 {
			ii++
			j--
			if s[ii-1] == j {
				return false
			}

		}
	}
	return true
}

func perm(s []int, index int, num *[100]int, ind *int) {
	if index == len(s)-1 {
		if checking(s) {
			d := 0
			*ind++
			for i := 7; i >= 0; i-- {
				d = d*10 + s[i]
			}
			num[*ind] = d

		}
	}
	for i := index; i < 8; i++ {
		swap(s, index, i)
		perm(s, index+1, num, ind)
		swap(s, index, i)
	}
}

func EightQueens() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	ss := [8]byte{}
	var num [100]int
	ind := 0
	perm(s, 0, &num, &ind)
	SortIntegerTable(num[:])
	for i := 8; i < 100; i++ {
		for j := 7; j >= 0; j-- {
			ss[j] = byte(num[i]%10 + 48)
			num[i] = num[i] / 10
		}

		for j := 0; j < 8; j++ {
			z01.PrintRune(rune(ss[j]))
		}
		z01.PrintRune(10)
	}
}
