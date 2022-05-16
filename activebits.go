package piscine

func ActiveBits(n int) int {
	cnt := 0
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	for n > 1 {
		if n%2 == 1 {
			cnt++
		}
		n = n / 2
	}
	return cnt + 1
}
