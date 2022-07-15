package main

import "fmt"

var myfunc func(int) int
var myInt int

func main(){
	myInt = 10

	myfunc = func(num int) int {
		fmt.Println("Inside anony")
		return num+1
	}

	test()
}

func test(){
	myfunc(10)
	fmt.Println(myInt)
	fmt.Println(myfunc)
}


/*
var myfunc func(int) int
myfunc is a variable name
myfunc ka type is func(int) int
myfunc is an anonymous function

*/