/*package main

import "github.com/01-edu/z01"

func main() {
	for i := 48; i <= 57; i++ {
		z01.PrintRune(rune(i))
	}
	z01.PrintRune(rune('\n'))
}
*/

package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.FindNextPrime(-2))
	fmt.Println(piscine.FindNextPrime(100))
}
