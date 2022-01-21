package main

import "fmt"
const myConstant = 30

const(
	
	// cat = iota //enumerated constants fromm 0
	cat int = iota + 1//enumerated constants fromm 1
	dog 
	camel 
	horse
)
func main() {
	// const myConstant = "hello" -> shadowing a constant
	// myConstant="there" -> not allowed
	fmt.Printf("Type of our constant is : %T\n",myConstant)
	fmt.Println("My constant is : ",myConstant)
	
	fmt.Println(cat,dog,camel,horse)

	// var num int32
	// fmt.Println(num+cat) error as num is int32 and cat is int-mixed match type
}