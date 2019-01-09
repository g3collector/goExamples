package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	beforeEcho1 := time.Now()
	Echo1()
	afterEcho1 := time.Now()
	afterEcho1.Sub(beforeEcho1)
	fmt.Println(afterEcho1)

	beforeEcho2 := time.Now()
	Echo2()
	afterEcho2 := time.Now()
	afterEcho2.Sub(beforeEcho2)
	fmt.Println(afterEcho2)

	beforeEcho3 := time.Now()
	Echo3()
	afterEcho3 := time.Now()
	afterEcho3.Sub(beforeEcho3)
	fmt.Println(afterEcho3)
}

func Echo1() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s = sep + os.Args[i]
		sep = " "
	}
	fmt.Println("Echo1", s)
}

func Echo2() {
	var s, sep string
	for _, v := range os.Args {
		s = sep + v
		sep = " "
	}
	fmt.Println("Echo2", s)
}

func Echo3() {
	s := strings.Join(os.Args, " ")
	fmt.Println("Echo3", s)
}
