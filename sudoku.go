package main

import (
	"fmt"
	"os"
)

func remove(s []int, index int) []int {

	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
func main() {
	arg := os.Args[1:]
	empty := 0
	var field [9][9]int
	if len(arg) != 9 {
		fmt.Println("Error")
		return
	}

	for i, line := range arg {
		if len(line) != 9 {
			fmt.Println("Error")
			return
		}
		for j, w := range line {
			if w >= 47 && w <= 58 {
				field[i][j] = int(w - '0')
			} else {
				empty++
			}

		}
	}
	for i := 0; i < 9; i++ { //checking squares
		for j := 0; j < 9; j++ {
			var starty, startx int = i / 3, j / 3 //Squares
			for y := starty * 3; y < (starty+1)*3; y++ {
				for x := startx * 3; x < (startx+1)*3; x++ {
					if y == i && x == j {
						continue
					}
					if field[y][x] != 0 {
						if field[y][x] == field[i][j] {
							fmt.Println("Error!")
							return
						}
					}
				}
			}
		}
	}
	for i := 0; i < 9; i++ { //checking lines
		for j := 0; j < 9; j++ {
			if field[i][j] != 0 {
				for k := 1; k < 9; k++ {
					if i+k < 9 && field[i][j] == field[i+k][j] {
						fmt.Println("Error")
						return
					}
					if k+j < 9 && field[i][j] == field[i][k+j] {
						fmt.Println("Error")
						return
					}
					if i-k >= 0 && field[i][j] == field[i-k][j] {
						fmt.Println("Error")
						return
					}
					if j-k >= 0 && field[i][j] == field[i][j-k] {
						fmt.Println("Error")
						return
					}
				}
			}

		}
	}

	elem := [81][]int{} //creating slice of possible values
	for i := 0; i < 81; i++ {
		for j := 0; j < 9; j++ {
			elem[i] = append(elem[i], j+1)
		}
		elem[i] = append(elem[i], -1)
	}

	for empty > 0 { //the solution
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if field[i][j] == 0 {
					temp := elem[i*9+j]
					for k := 1; k < 9; k++ { //removing numbers from slice  by LINES
						if k+j < 9 && 0 != field[i][k+j] {
							for h := 0; h < len(temp); h++ {
								if temp[h] == field[i][k+j] {
									temp = remove(temp, h)
									break
								}
							}

						}
						if i+k < 9 && 0 != field[i+k][j] {
							for h := 0; h < len(temp); h++ {
								if temp[h] == field[i+k][j] {
									temp = remove(temp, h)
									break
								}
							}

						}
						if j-k >= 0 && 0 != field[i][j-k] {
							for h := 0; h < len(temp); h++ {
								if temp[h] == field[i][j-k] {
									temp = remove(temp, h)
									break
								}
							}

						}
						if i-k >= 0 && 0 != field[i-k][j] {
							for h := 0; h < len(temp); h++ {
								if temp[h] == field[i-k][j] {
									temp = remove(temp, h)
									break
								}
							}

						}
					}
					var starty, startx int = i / 3, j / 3 //Squares

					for y := starty * 3; y < (starty+1)*3; y++ {
						for x := startx * 3; x < (startx+1)*3; x++ {
							if field[y][x] != 0 {
								for h := 0; h < len(temp); h++ {
									if temp[h] == field[y][x] {
										temp = remove(temp, h)
										break
									}
								}
							}
						}
					}
					if len(temp) == 2 {
						field[i][j] = temp[0]
						fmt.Println(i, j, field[i][j])
						empty--
					}
					elem[i*9+j] = temp
					// for m := 0; m < 81; m++ {
					// 	if m%8 == 0 {
					// 		fmt.Println()
					// 	}
					// 	fmt.Print(elem[m])
					// }
				}
			}
		}
	}

	fmt.Println(elem)

}
