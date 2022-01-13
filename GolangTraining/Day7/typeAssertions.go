package main

import "fmt"

func main() {
	var i interface{} = "hello" //empty interface can accept any type of data
	fmt.Println(i)
	i = 20
	fmt.Println(i)
	i = true
	fmt.Println(i)

	storeBool,OK := i.(string) //checks whether i is a string or not but since i is a bool so storeBool
	fmt.Println(storeBool,OK)

	i="Aishwarya"
	storeString,OK := i.(string) 
	fmt.Println(storeString,OK)

}

func checkType(anyType interface{}) { //empty interface with no methods or declarations

}