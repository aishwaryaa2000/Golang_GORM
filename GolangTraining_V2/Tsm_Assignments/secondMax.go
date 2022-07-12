package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	numberSlice := make([]int,0,10)
    var len,i uint64
	var userNumber int
	fmt.Print("Enter the count of numbers you wish to enter : ")
	_,err := fmt.Scanln(&len)

	if err!=nil{
		//If user enters a character i.e not a number and when user enters a negative number as len is uint
		fmt.Println("Either you entered a negative number or a character")
		log.Fatal(err)
	}

    for i=0;i<len;i++{
		fmt.Print("Enter number",i+1,": ")
		_,err2:=fmt.Scanln(&userNumber)
		if err2!=nil{
			//If user enters a character i.e not a number
			log.Fatal(err2)
		}
		numberSlice = append(numberSlice,userNumber)
	}

	secondLarge := secondMax(numberSlice)
	
	if secondLarge!=math.MinInt64{
	fmt.Println("Second largest number is : ",secondLarge)
	}else{
		fmt.Println("No second largest number")
	}
}

func secondMax(number []int) (int){

	secondLarge := math.MinInt64
	//Assign secondLarge as the most minimum value of int


	if len(number)<=1{
		return secondLarge
	}

	largest := number[0]
	for i:=1;i<len(number);i++{
		if number[i]>largest{
			secondLarge = largest
			largest=number[i]

		}else if number[i]<largest && number[i]>secondLarge{
			//Specify the else condition because in case of 3 3 3 there is no secondLarge element
			secondLarge = number[i]

		}
	}
	return secondLarge
}
