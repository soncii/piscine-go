package piscine

func FindNextPrime(nb int) int {
	if nb <= 0 {
		return 2
	}
	if nb == 1 {
		return 2
	}
	if nb == 2 {
		return 2
	}
	if nb == 3 {
		return 3
	}
	if nb == 4 {
		return 5
	}
	if nb == 5 {
		return 5
	}
	if nb%2 == 0 {
		return FindNextPrime(nb + 1)
	}
	for i := 3; i < nb/2; i += 2 {
		if nb%i == 0 {
			return FindNextPrime(nb + 2)
		}
	}
	return nb
}
