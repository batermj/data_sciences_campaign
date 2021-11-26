package main

import (
	"fmt"
	"time"
	"math"
)

const (
	e  = 2.71828182845904523536028747135266249775724709369995957496696763
	pi = 3.14159265358979323846264338327950288419716939937510582097494459
)

const (
	FlagUp Flags = 1 << iota // is up
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

const (
    B = 1 << (10 * iota)
    KiB // 1024
    MiB // 1048576
    GiB // 1073741824
    TiB // 1099511627776
    PiB // 1125899906842624
    EiB // 1152921504606846976
    ZiB // 1180591620717411303424
    YiB // 1208925819614629174706176
)

type Weekday int
type Flags uint

func main(){
	fmt.Println(e,pi)

	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
     
	fmt.Printf("%T %[1]v\n", noDelay)     // "time.Duration 0"
	fmt.Printf("%T %[1]v\n", timeout)     // "time.Duration 5m0s
	fmt.Printf("%T %[1]v\n", time.Minute)

	const ( 
		a=1
		b 
		c=2 
		d
	)
     
	fmt.Println(a, b, c, d)

	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Sunday,Monday,Tuesday,Wednesday,Thursday,Friday,Saturday)

	fmt.Println(FlagUp,FlagBroadcast,FlagLoopback,FlagPointToPoint,FlagMulticast)

	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 true"
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10000 false"
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))   // "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 true"
     
	fmt.Println(YiB/ZiB)

	var x float32 = math.Pi
	var y float64 = math.Pi
	var z complex128 = math.Pi
	fmt.Println(x,y,z)

	const Pi64 float64 = math.Pi
	var x2 float32 = float32(Pi64)
	var y2 float64 = Pi64
	var z2 complex128 = complex128(Pi64)
	fmt.Println(Pi64,x2,y2,z2)

	var f float64 = 212
	fmt.Println((f - 32) * 5 / 9)     // "100"; (f - 32) * 5 is a float64
	fmt.Println(5 / 9 * (f - 32))     // "0";   5/9 is an untyped integer, 0
	fmt.Println(5.0 / 9.0 * (f - 32))

	fmt.Printf("%T\n", 0)
	fmt.Printf("%T\n", 0.0)
	fmt.Printf("%T\n", 0i)
	fmt.Printf("%T\n", '\000') // "int32" (rune)

}

func IsUp(v Flags) bool     { return v&FlagUp == FlagUp }
func TurnDown(v *Flags)     { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }

