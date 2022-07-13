/* 
Maps or countArray[100]
*/

package main

import (
	"fmt"
	"log"
)

func main() {
	mySlice := make([]int,0,10)
    var len,i uint64
	userNumber:=0
	fmt.Print("Enter the count of numbers you wish to enter : ")
	_,err := fmt.Scanln(&len)

	if err!=nil || len==0{
		//If user enters a character i.e not a number and when user enters a negative number as len is uint
		fmt.Println("Either you entered a negative number or a character or a zero")
		log.Fatal(err)
	}


    fmt.Println("Please enter numbers greater than 0 and less than 100")
    for i=0;i<len;i++{
		fmt.Print("Enter number",i+1,": ")
		fmt.Scanln(&userNumber)
		if userNumber<=0 || userNumber >100{
			i--
			fmt.Println("Oops!This is an invalid number.Please enter a valid number.")
			continue
		} 
		mySlice = append(mySlice,userNumber)
	}

	getCount(mySlice)
}

func getCount(number []int) {

	/* var myMap = make(map[int]int)
	for _, value := range number {
		myMap[value]++
	}

	for i := 1; i < 100; i++ {
		if myMap[i] != 0 {
			fmt.Println("The number ",i," has occured ",myMap[i],"times")
		}
	}
	*/

	var countArray [100]int
    for _,value :=range number{
		countArray[value-1]++
	}
	for i:=0;i<100;i++{
		if countArray[i]!=0{
			fmt.Println("The number ",i+1,"has occured ",countArray[i]," times")
		}
	}
}