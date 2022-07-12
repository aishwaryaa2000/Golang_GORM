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

	SecondLarge,OkBool := secondMax(numberSlice)
	if OkBool{
	fmt.Println("Second largest number is : ",SecondLarge)
	}else{
		fmt.Println("No second largest number")
	}
}

func secondMax(number []int) (int,bool){

	
	if len(number)<=1{
		return 0,false
	}
	okBool := false
	secondLarge := math.MinInt64
	largest := number[0]
	for i:=1;i<len(number);i++{
		if number[i]>largest{
			largest=number[i]
			secondLarge = largest
			okBool = true

		}else if number[i]<largest && number[i]>secondLarge{
			secondLarge = number[i]
			okBool = true

		}
	}
	return secondLarge,okBool
}
