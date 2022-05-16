package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	for i := 0; i < len(arg); i++ {
		if arg[i] == "01" || arg[i] == "galaxy 01" || arg[i] == "galaxy" {
			fmt.Println("Alert!!!")
			break
		}
	}
}
