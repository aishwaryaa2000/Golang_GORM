// a state in which program doesn't know what to do next
// the execuetion stops as soon as panic occurs

package main

import "fmt"

func main(){
	fmt.Println("Opening a resource")
	defer func() {
		//this is anonymous function
		fmt.Println("Closing a resource")
		details := recover() //it gives details about the panic
		fmt.Println("Details : ",details) //details me "Panicking don't know what to do"
	}()
	panic("Panicking don't know what to do")
//even if panic occurs,defer func() is called
}

/*
opening a resource
Closing a resource
Details : Panicking don't know what to do
*/