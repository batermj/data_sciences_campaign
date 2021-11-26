package main

import (
	"fmt"
)


func main(){
	var runes []rune
	for _, r := range "Hello, BF" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes) // "['H' 'e' 'l' 'l' 'o' ',' ' ' 'B' 'F']"

}

func appendInt(x []int, y int) []int {
         
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
    	z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	z[len(x)] = y
	return z
}

