package main

import (
	"fmt"
	"os"
)

// Main print out the file path with also args
func main() {
	var s, sep string

	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
}
