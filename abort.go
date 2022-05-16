package piscine

func Abort(a, b, c, d, e int) int {
	table := []int{a, b, c, d, e}
	len := len(table) - 1
	for i := 0; i < len; i++ {
		for j := 0; j < len-i; j++ {
			if table[j] > table[j+1] {
				temp := table[j]
				table[j] = table[j+1]
				table[j+1] = temp
			}
		}
	}
	return table[2]
}
