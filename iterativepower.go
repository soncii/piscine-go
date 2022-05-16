package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	if power == 1 {
		return nb
	}
	sum := 1
	for i := 0; i < power; i++ {
		sum *= nb
	}
	return sum
}
