// a state in which program doesn't know what to do next
// the execuetion stops as soon as panic occurs

package main

import "fmt"

func main(){
	fmt.Println("Opening a resource")
	defer func() {
		fmt.Println("Closing a resource")
		details := recover()
		fmt.Println("Details : ",details)
	}()
	panic("Panicking don't know what to do")
//even if panic occurs,defer func() is called
}