package main

import "fmt"

func main(){
	var p = f()
	fmt.Printf("%d\n",(int)(*p))
}

func f() *int {
	v := 1
	return &v 
}

