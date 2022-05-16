package piscine

func AlphaCount(s string) int {
	a := []rune(s)
	cnt := 0
	for i := 0; i < len(a); i++ {
		if (a[i] > 64 && a[i] < 91) || (a[i] > 96 && a[i] < 123) {
			cnt++
		}
	}
	return cnt
}
