package main

import(
	"fmt"
	"reflect"
)

func main(){
	var myInt int
	var myString string
	var myArray = [3]int{}
	var mySlice = []int{}


	fmt.Println("Zero value of int is : ",myInt)  //0
	fmt.Println("Zero value of string is : ",myString) // 
	fmt.Println("Zero value of array is : ",myArray)  // [0 0 0]
	fmt.Println(reflect.TypeOf(myArray))  // [3]int
	fmt.Println("Zero value of slice is : ",mySlice)  // []
	fmt.Println(reflect.TypeOf(mySlice))  // []int

}

