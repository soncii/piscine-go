package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	sum := 1
	for i := 1; i <= nb; i++ {
		sum = sum * i
		if sum < 0 {
			return 0
		}
	}
	return sum
}
