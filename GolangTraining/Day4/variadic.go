package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(sum(""))
	fmt.Println(sum("",10))
	fmt.Println(sum("",10,20,30,40))

	var mySlice = []int{10,20,30,40}
	fmt.Println(reflect.TypeOf(mySlice))
	fmt.Println(sum("",mySlice...))
	// mySlice has variable arguments inside it or many values inside it so three dots... while func call



}

func sum(str string,num...int) int{
	total := 0
	for _,value := range num{
		total += value
	}

	return total
}

// variable number of arguments 
// first we need to type the compulsory parameters and then the variadic variables with ...