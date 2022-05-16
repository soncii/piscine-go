package piscine

func Map(f func(int) bool, a []int) []bool {
	res := make([]bool, len(a))
	for i, w := range a {
		res[i] = f(w)
	}
	return res
}
