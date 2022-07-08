package main

import (
	"fmt"
)

func main() {
	var no int16 = 10
	var nooo int
	nooo = int(no)
	fmt.Println(no, nooo)//10 10

	var no2 int16 = 10000
	//binary equivalent is 1 0 0 1 1 1 0 0 0 1 0 0 0 0
	var nooo2 int8
	nooo2 = int8(no2)
	//binary equivalent is 0 0 0 1 0 0 0 0 i.e the last 8 bits or LSB 8 bits
	//0 0 0 1 0 0 0 0 decimal equivalent is 16
	fmt.Println(no2, nooo2) //10000 16

}
