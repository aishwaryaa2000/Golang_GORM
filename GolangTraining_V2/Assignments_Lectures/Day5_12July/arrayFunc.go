package main

import "fmt"

func main() {
	var myArray = [3]int{10, 20, 30}
	fmt.Println("Value is in main is : ",myArray)
	fmt.Println("Address in main is : ",&myArray[0])


	test1(myArray)
	fmt.Println("Value is in main after test1 is called : ",myArray)

	 test2(&(myArray))
	 fmt.Println("Value is in main after test2 is called: ",myArray)


}

func test1(myArray [3]int) {
	myArray[0] = 100
	fmt.Println("Address in test1 is : ",&myArray[0])
	fmt.Println("Value is in test1 is : ",myArray)

}

func test2(myArray *[3]int){
	myArray[1]=200//no need to dereference it
	fmt.Println("Address in test2 is : ",&myArray[0])
	fmt.Println("Value is in test2 is : ",myArray)

}

/* 
Value is in main is :  [10 20 30]
Address in main is :  0xc42000a1e0

Address in test1 is :  0xc42000a240
Value is in test1 is :  [100 20 30]
Value is in main after test1 is called :  [10 20 30]

Address in test2 is :  0xc42000a1e0
Value is in test2 is :  &[10 200 30]
Value is in main after test2 is called:  [10 200 30]

*/