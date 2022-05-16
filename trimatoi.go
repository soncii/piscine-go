package piscine

func TrimAtoi(s string) int {
	a := []rune(s)
	j := 0
	var num []int
	sign := true
	minus := false
	for i := 0; i < len(a); i++ {
		if sign && a[i] == '-' {
			minus = true
			sign = false
		}
		if a[i] >= 48 && a[i] <= 57 {
			sign = false
			num = append(num, int(a[i]-'0'))
			j++
		}
	}
	n := 0
	for i := 0; i < j; i++ {
		n = n * 10
		n = n + num[i]
	}
	if minus {
		return n * -1
	}
	return n
}
