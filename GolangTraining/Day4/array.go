package main

import "fmt"

func main() {
	var myArray = [3]int{10, 20, 30}

	// 	var myArray = [...]int{10, 20} 
	//  ... automatically declares the size of array acc to the no. of elements present ,above case size = 2
	
	fmt.Println(myArray)

	// array is call by value,so if we assign it to another variable then A COPY OF THE ARRAY is saved
	//if we change anything in the COPY OF ARRAY then ORIGINAL ARRAY WON'T CHANGE
	
	var copyMyArray = myArray
	copyMyArray[0] = 40
	fmt.Println("My array is : ",myArray)
	fmt.Println("Copy of array : ",copyMyArray)

}