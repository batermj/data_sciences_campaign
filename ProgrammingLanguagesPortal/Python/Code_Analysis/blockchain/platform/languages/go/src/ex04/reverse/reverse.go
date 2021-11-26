package main

import (
	"fmt"
)

func main(){
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a) // "[5 4 3 2 1 0]"
	fmt.Printf("%T\n",a)

	months := [...]string{1: "January", 2: "Febrary", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "Auguest", 9: "September", 10: "October", 11: "November", 12: "December"}
	fmt.Println(months)
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)     
	fmt.Println(summer) 

	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			} 
		}
	}


}
     
// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	} 
}

