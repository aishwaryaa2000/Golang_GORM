package main

import "fmt"
// import "reflect"

func main(){
	fmt.Println(test2())
	fnreturned:= test2()
	fnreturned() 
    //this function is getting invoked

	// test2()() this can also be directly used instead of above two statements


	// fmt.Println(test1) prints address of test1 -example 0xe8c660
	// fmt.Println(reflect.TypeOf(test1)) returns func() 
	
}

func test1(){
	fmt.Println("Inside test1")
}

func test2() func(){
	fmt.Println("Inside test2")
	return test1 //this is the address of test1
}