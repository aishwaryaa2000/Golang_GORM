// defer has to be followed by func call
package main

import "fmt"

func main(){
	number := 50
	fmt.Println("Start of main")
	f1() 
	defer f2(number) 
	//stored in a stack  f1()->f2() i.e LIFO so f2() executed first

	number=100
	fmt.Println("Value of number in main : ",number)
	fmt.Println("End of main")

}

func f1(){
	defer fmt.Println("Defer inside f1")
	fmt.Println("End of f1")
	// as soon as execuetion is at "End of f1" then defer fmt.println is called
}


func f2(n int){
	fmt.Println("End of f2")
	fmt.Println("Value of number in f2 is : ",n)
}

/*defer func() is stored in a stack
will defer or avoid f1()'s execuetion until surrounding func's i.e main() ends



open resource
defer close() 
panic
after panic defer close() is executed so it closes the opened resource
*/
