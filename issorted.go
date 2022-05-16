package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	i := 0
	pattern := 0
	for i = 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) != 0 {
			pattern = f(a[i], a[i+1])
			break
		}
	}
	if i == len(a)-1 {
		return true
	}
	if pattern < 0 {
		for i := 0; i < len(a)-1; i++ {
			if f(a[i], a[i+1]) > 0 {
				return false
			}
		}
		return true
	}

	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) < 0 {
			return false
		}
	}
	return true
}
