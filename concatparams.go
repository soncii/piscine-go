package piscine

func ConcatParams(args []string) string {
	var s string
	for i := 0; i < len(args); i++ {
		s = s + args[i]
		if i != len(args)-1 {
			s = s + "\n"
		}
	}
	return s
}
