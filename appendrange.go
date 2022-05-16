package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}
	n := []int{}
	for i := min; i < max; i++ {
		n = append(n, i)
	}
	return n
}
