package main

import (
	"fmt"
	"os"
	"strings"
	"ex01/stringutil"
)

func main() {
	fmt.Printf("Hello, world.\n")
	fmt.Printf(stringutil.Reverse("!oG ,olleH"))
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " " 
	}
	fmt.Println(s)

	fmt.Println("-----")
	s,sep = "",""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " " 
	}
	fmt.Println(s)

	// effectively join string
	fmt.Println(strings.Join(os.Args[1:], " "))

	// fmt in charge of the output format for args
	fmt.Println(os.Args[1:])
}
