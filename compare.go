package piscine

func Compare(a, b string) int {
	if a > b {
		return 1
	}
	if a < b {
		return -1
	}

	return 0

	// s1 := []rune(a)
	// s2 := []rune(b)
	// l := 0
	// if len(s1) >= len(s2) {
	// 	l = len(s1) - 1
	// } else {
	// 	l = len(s2) - 1
	// }
	// for i := 0; i < l-1; i++ {
	// 	if s1[i] > s2[i] {
	// 		return 1
	// 	} else if s1[i] < s2[i] {
	// 		return -1
	// 	}
	// }
	// if len(s1) > len(s2) {
	// 	return 1
	// }
	// if len(s1) < len(s2) {
	// 	return -1
	// }
	// return 0
}
