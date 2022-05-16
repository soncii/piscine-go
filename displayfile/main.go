package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("File name missing")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	name := os.Args[1]
	file, err := os.Open(name)
	fmt.Println(err)
	if err != nil {
		fmt.Println("File name missing")
		return
	}

	s := make([]byte, 200)
	n1, err := file.Read(s)
	fmt.Printf("The length of the string: %v \n%v \n%v \n%v\n \n The error: %s\n string: %s\n\n\n\n", n1, 15324, 564, 45671, err, string(s))
	fmt.Print(string(s[:n1]))
	if err != nil {
		fmt.Println("File name missing")
		return
	}
}
