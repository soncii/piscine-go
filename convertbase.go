package piscine

func printNbrBase(nbr int, base string) string {
	ind := 0
	b := []rune(base)
	l := len(b)
	s := []rune{}
	d := 0
	for nbr >= l {
		d++
		if ind == 1 {
			s = append(s, b[nbr%l+1])
		}
		s = append(s, b[nbr%l])
		nbr = nbr / l
	}
	s = append(s, b[nbr])
	n := make([]rune, d+2)
	for i := d; i >= 0; i-- {
		n[d-i] = s[i]
	}
	return string(n)
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	n := AtoiBase(nbr, baseFrom)
	r1 := []rune(printNbrBase(n, baseTo))
	l := len(r1) - 1
	if r1[l] == '\x00' {
		res := make([]rune, len(r1)-1)
		for i := 0; i < len(res); i++ {
			res[i] = r1[i]
		}
		return string(res)
	} else {
		return printNbrBase(n, baseTo)
	}
}
