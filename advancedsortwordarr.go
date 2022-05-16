package piscine

// func Compa(a, b string) int {
// 	s1 := []rune(a)
// 	s2 := []rune(b)
// 	l := 0
// 	if len(s1) >= len(s2) {
// 		l = len(s1) - 1
// 	} else {
// 		l = len(s2) - 1
// 	}
// 	for i := 0; i < l-1; i++ {
// 		if s1[i] > s2[i] {
// 			return 1
// 		} else if s1[i] < s2[i] {
// 			return -1
// 		}
// 	}
// 	if len(s1) > len(s2) {
// 		return 1
// 	}
// 	if len(s1) < len(s2) {
// 		return -1
// 	}
// 	return 0
// }

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	len := len(a) - 1
	var temp string
	for i := 0; i < len; i++ {
		for j := 0; j < len-i; j++ {
			if Compare(a[j], a[j+1]) == 1 {
				temp = a[j]
				a[j] = a[j+1]
				a[j+1] = temp
			}
		}
	}
}
