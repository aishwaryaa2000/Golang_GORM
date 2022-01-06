package main

import "fmt"

func main(){
	fnreturned:= test2()
	fnreturned() 
	//this function is getting invoked
}

func test1(){
	fmt.Println("Inside test1")
}

func test2() func(){
	fmt.Println("Inside test2")
	return test1
}