package piscine

func Sqrt(nb int) int {
	i := 1
	sq := 1
	if nb < 1 {
		return 0
	}
	for sq <= nb {
		sq = i * i
		if sq == nb {
			return i
		}
		i++
	}
	return 0
}
