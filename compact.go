package piscine

func Compact(ptr *[]string) int {
	cnt := 0
	s := *ptr
	for i := 0; i < len(s); i++ {
		if s[i] != "" {
			s[cnt] = s[i]
			cnt++
		}
	}
	*ptr = s[0:cnt]
	return cnt
}
