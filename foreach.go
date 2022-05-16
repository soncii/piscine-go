package piscine

func ForEach(f func(int), a []int) {
	for _, w := range a {
		f(w)
	}
}
