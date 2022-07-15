package main

import "fmt"

func main(){
	func(){
		fmt.Println("Inside anonymous func")
	}() //Immediately invoked function expression - IIFE

	anonyFunc()

}

func anonyFunc(){
	func(s string){
		fmt.Println(s,"Inside non main anonymous func")
	}("Hello Aishwarya")
}


func anonyFunc1(){
	var number = func(s string) int { //number will have return value of the funci.e int
		fmt.Println(s,"Inside non main anonymous func")
		return 100
	}("Hello Aayush")

	fmt.Println("Return value is : ",number)
}
