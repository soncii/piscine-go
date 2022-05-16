package piscine

func Rot14(s string) string {
	str := []byte(s)
	for i := 0; i < len(str); i++ {
		if str[i] >= 'A' && str[i] <= 'L' {
			str[i] = byte(int(str[i]) + 14)
			continue
		}
		if str[i] >= 'a' && str[i] <= 'l' {
			str[i] = byte(int(str[i]) + 14)
			continue
		}
		if str[i] >= 'L' && str[i] <= 'Z' {
			str[i] = byte(int(str[i]) + 14 - 26)
			continue

		}
		if str[i] >= 'l' && str[i] <= 'z' {
			str[i] = byte(int(str[i]) + 14 - 26)
			continue
		}
	}
	return string(str)
}
