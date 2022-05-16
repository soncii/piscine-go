package piscine

func SortWordArr(a []string) {
	len := len(a) - 1
	var temp string
	for i := 0; i < len; i++ {
		for j := 0; j < len-i; j++ {
			if a[j] > a[j+1] {
				temp = a[j]
				a[j] = a[j+1]
				a[j+1] = temp
			}
		}
	}
}
