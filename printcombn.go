package piscine

import (
	"github.com/01-edu/z01"
)

func PrintCombN(n int) {
	if n == 1 {
		for i := '0'; i <= '9'; i++ {
			z01.PrintRune(rune(i))
			if i != '9' {
				z01.PrintRune(rune(','))
				z01.PrintRune(rune(' '))
			}
		}
	}
	if n == 2 {
		for i := '0'; i <= '8'; i++ {
			for j := i + 1; j <= '9'; j++ {
				z01.PrintRune(rune(i))
				z01.PrintRune(rune(j))
				if i == '8' && j == '9' {
				} else {
					z01.PrintRune(rune(','))
					z01.PrintRune(rune(' '))
				}
			}
		}
	}
	if n == 3 {
		for i := '0'; i <= '7'; i++ {
			for j := i + 1; j <= '8'; j++ {
				for k := j + 1; k <= '9'; k++ {
					z01.PrintRune(rune(i))
					z01.PrintRune(rune(j))
					z01.PrintRune(rune(k))
					if i == '7' && j == '8' && k == '9' {
					} else {
						z01.PrintRune(rune(','))
						z01.PrintRune(rune(' '))
					}
				}
			}
		}
	}
	if n == 4 {
		for l := '0'; l <= '6'; l++ {
			for i := l + 1; i <= '7'; i++ {
				for j := i + 1; j <= '8'; j++ {
					for k := j + 1; k <= '9'; k++ {
						z01.PrintRune(rune(l))
						z01.PrintRune(rune(i))
						z01.PrintRune(rune(j))
						z01.PrintRune(rune(k))
						if l == '6' && i == '7' && j == '8' && k == '9' {
						} else {
							z01.PrintRune(rune(','))
							z01.PrintRune(rune(' '))
						}
					}
				}
			}
		}
	}
	if n == 5 {
		for o := '0'; o <= '5'; o++ {
			for l := o + 1; l <= '6'; l++ {
				for i := l + 1; i <= '7'; i++ {
					for j := i + 1; j <= '8'; j++ {
						for k := j + 1; k <= '9'; k++ {
							z01.PrintRune(rune(o))
							z01.PrintRune(rune(l))
							z01.PrintRune(rune(i))
							z01.PrintRune(rune(j))
							z01.PrintRune(rune(k))
							if o == '5' && l == '6' && i == '7' && j == '8' && k == '9' {
							} else {
								z01.PrintRune(rune(','))
								z01.PrintRune(rune(' '))
							}
						}
					}
				}
			}
		}
	}
	if n == 6 {
		for u := '0'; u <= '4'; u++ {
			for o := u + 1; o <= '5'; o++ {
				for l := o + 1; l <= '6'; l++ {
					for i := l + 1; i <= '7'; i++ {
						for j := i + 1; j <= '8'; j++ {
							for k := j + 1; k <= '9'; k++ {
								z01.PrintRune(rune(u))
								z01.PrintRune(rune(o))
								z01.PrintRune(rune(l))
								z01.PrintRune(rune(i))
								z01.PrintRune(rune(j))
								z01.PrintRune(rune(k))
								if u == '4' && o == '5' && l == '6' && i == '7' && j == '8' && k == '9' {
								} else {
									z01.PrintRune(rune(','))
									z01.PrintRune(rune(' '))
								}
							}
						}
					}
				}
			}
		}
	}
	if n == 7 {
		for y := '0'; y <= '3'; y++ {
			for u := y + 1; u <= '4'; u++ {
				for o := u + 1; o <= '5'; o++ {
					for l := o + 1; l <= '6'; l++ {
						for i := l + 1; i <= '7'; i++ {
							for j := i + 1; j <= '8'; j++ {
								for k := j + 1; k <= '9'; k++ {
									z01.PrintRune(rune(y))
									z01.PrintRune(rune(u))
									z01.PrintRune(rune(o))
									z01.PrintRune(rune(l))
									z01.PrintRune(rune(i))
									z01.PrintRune(rune(j))
									z01.PrintRune(rune(k))
									if y == '3' && u == '4' && o == '5' && l == '6' && i == '7' && j == '8' && k == '9' {
									} else {
										z01.PrintRune(rune(','))
										z01.PrintRune(rune(' '))
									}
								}
							}
						}
					}
				}
			}
		}
	}
	if n == 8 {
		for t := '0'; t <= '2'; t++ {
			for y := t + 1; y <= '3'; y++ {
				for u := y + 1; u <= '4'; u++ {
					for o := u + 1; o <= '5'; o++ {
						for l := o + 1; l <= '6'; l++ {
							for i := l + 1; i <= '7'; i++ {
								for j := i + 1; j <= '8'; j++ {
									for k := j + 1; k <= '9'; k++ {
										z01.PrintRune(rune(t))
										z01.PrintRune(rune(y))
										z01.PrintRune(rune(u))
										z01.PrintRune(rune(o))
										z01.PrintRune(rune(l))
										z01.PrintRune(rune(i))
										z01.PrintRune(rune(j))
										z01.PrintRune(rune(k))
										if t == '2' && y == '3' && u == '4' && o == '5' && l == '6' && i == '7' && j == '8' && k == '9' {
										} else {
											z01.PrintRune(rune(','))
											z01.PrintRune(rune(' '))
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	if n == 9 {
		for r := '0'; r <= '1'; r++ {
			for t := r + 1; t <= '2'; t++ {
				for y := t + 1; y <= '3'; y++ {
					for u := y + 1; u <= '4'; u++ {
						for o := u + 1; o <= '5'; o++ {
							for l := o + 1; l <= '6'; l++ {
								for i := l + 1; i <= '7'; i++ {
									for j := i + 1; j <= '8'; j++ {
										for k := j + 1; k <= '9'; k++ {
											z01.PrintRune(rune(r))
											z01.PrintRune(rune(t))
											z01.PrintRune(rune(y))
											z01.PrintRune(rune(u))
											z01.PrintRune(rune(o))
											z01.PrintRune(rune(l))
											z01.PrintRune(rune(i))
											z01.PrintRune(rune(j))
											z01.PrintRune(rune(k))
											if r == '1' && t == '2' && y == '3' && u == '4' && o == '5' && l == '6' && i == '7' && j == '8' && k == '9' {
											} else {
												z01.PrintRune(rune(','))
												z01.PrintRune(rune(' '))
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
