package piscine

func Max(a []int) int {
	m := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] > m {
			m = a[i]
		}
	}
	return m
}
