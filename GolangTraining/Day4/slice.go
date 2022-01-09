package main

import "fmt"

func main() {
	mySlice := make([]int, 5)
	// len,capacity where capacity >= len,by default length = capacity

	fmt.Println(mySlice)
	for i:=0 ;i<len(mySlice) ;i++{
		mySlice[i]=(i+1)*10
	}

	fmt.Println("Contents inside mySlice before any modification : ",mySlice)

	//call by reference, any changes in copyMyslice is reflected in the Myslice
	copyMySlice := mySlice
	copyMySlice[0] =100
	fmt.Println("Contents inside copyMySlice after modifying copyMySlice : ",copyMySlice)
	fmt.Println("Contents inside mySlice after modifying copyMySlice : ",mySlice)
	fmt.Println("Length of mySlice : ",len(mySlice))
	fmt.Println("Capacity of mySlice : ",cap(mySlice))


	// mySlice = append(mySlice, 200,300,400)
	// fmt.Println("Contents inside mySlice is : ",mySlice)
	// fmt.Println("Length of mySlice : ",len(mySlice))
	// fmt.Println("Capacity of mySlice : ",cap(mySlice))


}

// dynamic data structure
