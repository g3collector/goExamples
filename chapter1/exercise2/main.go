package main

import (
	"fmt"
	"os"
)

/**
Print the index of the range and also the argument
*/
func main() {
	var s, sep string
	for i, v := range os.Args[0:] {
		sep = " "
		s += string(i) + sep + v + sep
	}

	fmt.Println(s)
}
