package tempconv

/*
 * Life time of variables
 * 
 *
*/

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func main(){
	//medals := []string{"gold", "silver", "bronze"}
	///fmt.Println(medals)

	fmt.Println("Neo\n")

	fmt.Println("%g\n", BoilingC-FreezingC) // "100" °C
	boilingF := CToF(BoilingC)
	fmt.Println("%g\n", boilingF-CToF(FreezingC)) // "180" °F
	//fmt.Printf("%g\n", boilingF-FreezingC)       // compile error: type mismatch

	c := FToC(212.0)
	fmt.Println(c.String()) // "100°C"
	fmt.Printf("%v\n", c)
	fmt.Printf("%s\n", c)
	fmt.Println(c)
	fmt.Printf("%g\n", c)
	fmt.Println(float64(c)) // "100"; does not call String

}
