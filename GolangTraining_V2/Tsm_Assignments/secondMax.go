package main

import (
	"fmt"
	"log"
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

func secondMax(number []int) (int64,bool){

	
	if len(number)<=1{
		return 0,false
	}
}
