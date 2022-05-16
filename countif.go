package piscine

func CountIf(f func(string) bool, tab []string) int {
	res := 0
	for _, w := range tab {
		if f(w) {
			res++
		}
	}
	return res
}
