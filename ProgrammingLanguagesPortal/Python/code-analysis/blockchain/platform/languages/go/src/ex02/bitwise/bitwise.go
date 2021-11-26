package main

import (
	"fmt"
	"math"
)

func main(){
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
 
	var apples int32 = 1
	var oranges int16 = 2
	///var compote int = apples + oranges // compile error
	var compote = int(apples) + int(oranges)
	fmt.Println(apples,oranges,compote)

	fmt.Printf("%08b\n", x)    // "00100010", the set {1, 5}
	fmt.Printf("%08b\n", y)    // "00000110", the set {1, 2}
	fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
	fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
	fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
	fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}
	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i) // "1", "5"
		} 
	}
     
	fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}

	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
	x64 := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x64)

	
	ascii := 'a'
	unicode01 := '言'
	unicode02 := 'қ'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii) // "97 a 'a'" 
	fmt.Printf("%d %[1]c %[1]q\n", unicode01) 
	fmt.Printf("%d %[1]c %[1]q\n", unicode02)  
	fmt.Printf("%d %[1]q\n", newline) // "10 '\n'"

	
	var f float32 = 16777216 // 1 << 24
	fmt.Println(f == f+1)    // "true"!
	fmt.Println(math.MaxFloat32,math.MaxFloat64)

     
	const Avogadro = 6.02214129e23
	const Planck   = 6.62606957e-34
	fmt.Println(Avogadro,Planck)


	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d exp(x) = %8.3f\n", x, math.Exp(float64(x)))
	}

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) //  "0 -0 +Inf -Inf NaN"

	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan) // "false false false"

}
