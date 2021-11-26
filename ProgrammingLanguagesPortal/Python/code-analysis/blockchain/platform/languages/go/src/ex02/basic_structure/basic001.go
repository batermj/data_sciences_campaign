package main

import (
	"fmt"
)

func main(){
	p := new(int)   // p, of type *int, points to an unnamed int variable
	fmt.Println(*p) // "0"
	*p = 2          // sets the unnamed int to 2
	fmt.Println(*p) // "2"

	p2 := new(int)
	q2 := new(int)
	fmt.Println(p2 == q2) // "false"

}
