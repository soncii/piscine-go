package piscine

func SortIntegerTable(table []int) {
	len := len(table) - 1
	var temp int
	for i := 0; i < len; i++ {
		for j := 0; j < len-i; j++ {
			if table[j] > table[j+1] {
				temp = table[j]
				table[j] = table[j+1]
				table[j+1] = temp
			}
		}
	}
}
