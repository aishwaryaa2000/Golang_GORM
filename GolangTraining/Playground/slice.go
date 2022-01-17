package main

import (
	"fmt"
)

var s = make([]int, 7, 8)


func main() {
	for i := 0; i < 7; i++ {
		s[i] = i
	}
	a := append(s,10)
	fmt.Println(a)
	fmt.Println(s)

	// fmt.Printf("Inside main the address of a is %p", &a)
	// fmt.Printf("\nAddr of first element inside main: %p\n", &a[0])

}

// func Test(ss *[]int) {
// 	// a = append(ss, 100)
// 	fmt.Println(ss)
// 	fmt.Printf("Inside Test the address of a is %p", &ss)
// 	// fmt.Printf("\nAddr of first element inside Test: %p\n", &ss[0])

// }