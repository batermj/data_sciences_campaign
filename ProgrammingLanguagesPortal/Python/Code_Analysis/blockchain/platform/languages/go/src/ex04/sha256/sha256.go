package main

import (
	"fmt"
	"crypto/sha256"
)

func main() {
	testArray()

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}

func testArray(){
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])
	fmt.Printf("%T\n",a)
	
	for i,v := range a {
		fmt.Printf("%d,%d\n",i,v)
	}

	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(q)
	fmt.Println(r)

	q2 := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q2)

	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "¥"}
	fmt.Println(RMB, symbol[RMB])
	for i,v := range symbol{
		fmt.Printf("%d %s\n",i,v)
	}

	r2	 := [...]int{99: -1}
	fmt.Printf("%T\n",r2)
	fmt.Println(r2)


	a2 := [2]int{1, 2}
	b2 := [...]int{1, 2}
	c2 := [2]int{1, 3}
	fmt.Println(a2 == b2, a2 == c2, b2 == c2) // "true false false"

	d := [3]int{1, 2}
	fmt.Println(a == d)

}
