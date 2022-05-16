package piscine

func Unmatch(a []int) int {
	c := 0
	for _, s := range a {
		c = 0
		for _, ss := range a {
			if s == ss {
				c++
			}
		}

		if c%2 != 0 {
			return s
		}
	}
	return -1
}
