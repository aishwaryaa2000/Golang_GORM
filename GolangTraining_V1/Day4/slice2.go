package main

import "fmt"

func main() {
	// var mySlice []int
	// fmt.Println(mySlice)   ->  []
	// fmt.Println(mySlice==nil)    -> true

	mySlice := make([]int,5,5)
	fmt.Println(mySlice)
	for i:=0 ;i<len(mySlice) ;i++{
		mySlice[i]=(i+1)*10
	}
	fmt.Println("\nContents in mySlice : ",mySlice)
	fmt.Println("Length of mySlice : ",len(mySlice))
	fmt.Println("Capacity of mySlice : ",cap(mySlice))
	fmt.Println("Address of first elemet of mySlice : ",&mySlice[0])

	// if number of items appended is less than capacity then capacity doubles

	// if we append outside capacity then address location changes.
	//  current capacity = 5 and current length = 5.Let's append 2 elements so address of mySlice will change

   mySlice = append(mySlice, 60,70)
   fmt.Println("\nContents in mySlice : ",mySlice)
   fmt.Println("Length of mySlice : ",len(mySlice))
   fmt.Println("Capacity of mySlice : ",cap(mySlice))
   fmt.Println("Address of first elemet of mySlice : ",&mySlice[0])

   //new capacity = 10 and length till now is 7. So let's append one element  
   //if we append inside new capacity then address location remains same 
  
   mySlice = append(mySlice, 80,90)
   fmt.Println("\nContents in mySlice : ",mySlice)
   fmt.Println("Length of mySlice : ",len(mySlice))
   fmt.Println("Capacity of mySlice : ",cap(mySlice))
   fmt.Println("Address of first elemet of mySlice : ",&mySlice[0])

    //Capacity till now=10 and length till now=9. So,let's add one element
    //if we append inside capacity then address location remains same 
	//So,address location of &copyMySlice[0] and &mySlice[0] is same
   copyMySlice := append(mySlice,100)
   fmt.Println("\nContents in copyMySlice : ",copyMySlice)
   fmt.Println("Length of copyMySlice : ",len(copyMySlice))
   fmt.Println("Capacity of copyMySlice : ",cap(copyMySlice))
   fmt.Println("Address of first elemet of copyMySlice : ",&copyMySlice[0])
   fmt.Println("Address of first elemet of MySlice : ",&mySlice[0])

   //If we append to copyMySlice then it won't be reflected in MySlice.
   fmt.Println("\nContents in mySlice : ",mySlice)

    //Capacity till now=10 and length till now=9. So,let's add three element
    //if we append oustide capacity then address location will change  
	//So,address location of &copyMySlice[0] and &mySlice[0] are different
   copyMySlice = append(mySlice,100,110,120)
   fmt.Println("\nContents in copyMySlice : ",copyMySlice)
   fmt.Println("Length of copyMySlice : ",len(copyMySlice))
   fmt.Println("Capacity of copyMySlice : ",cap(copyMySlice))
   fmt.Println("Address of first elemet of copyMySlice : ",&copyMySlice[0])
   fmt.Println("Address of first elemet of MySlice : ",&mySlice[0])

   //will print same address
   mySlice = copyMySlice
   fmt.Println("\nAddress of first elemet of copyMySlice : ",&copyMySlice[0])
   fmt.Println("Address of first elemet of MySlice : ",&mySlice[0])



}

// C:\xampp\htdocs\Swabhav_Techlabs\GolangTraining\Day4>go run slice.go
// Contents inside mySlice is :  [10 20 30 40 50]
// Length of mySlice :  5
// Capacity of mySlice :  5
// Contents inside mySlice is :  [10 20 30 40 50 200 300 400]
// Length of mySlice :  8
// Capacity of mySlice :  10

// C:\xampp\htdocs\Swabhav_Techlabs\GolangTraining\Day4>go run slice.go
// Contents inside mySlice is :  [10 20 30 40 50]
// Length of mySlice :  5
// Capacity of mySlice :  5
// Contents inside mySlice is :  [10 20 30 40 50 200 300 400 500 600 700 800 1000 9000]
// Length of mySlice :  14
// Capacity of mySlice :  14



