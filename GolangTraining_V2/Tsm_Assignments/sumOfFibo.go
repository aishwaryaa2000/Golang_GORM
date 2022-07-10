package main

import (
	"errors"
	"fmt"
	"log"
)

func sumOfFibonaci(noOfTerms int) (uint64,error) {
	//Fibonacci series : 0 1 1 2 3 8...
	if noOfTerms<0{
		err := errors.New("Number of terms cannot be less than zero")
		return 0,err
	}
	if noOfTerms==1 || noOfTerms==0 {
		return 0,nil
	}else{
		var(
			//Taking uint64 because we know that the numbers will always be positive thereby increasing the range
			sumTotal, firstNum, secondNum,thirdNum uint64= 1,0,1,0
		  )

		fmt.Print("Fibonacci series : ",firstNum," ",secondNum," ")
		for i:=0;i<noOfTerms-2;i++{
			thirdNum = firstNum+secondNum
			fmt.Print(thirdNum," ")
			sumTotal += thirdNum
			firstNum=secondNum
			secondNum=thirdNum
		}

		return sumTotal,nil

	}
	
}
func main(){
	var no int 
   
    fmt.Print("Enter an integer : ")
	_, err := fmt.Scanf("%d", &no)
	if err!=nil{
		//If user doesn't enter a number then simply raise an error and exit
		//Example-If user enters an alphabet then raise an 
		log.Fatal(err)
	}	
	sum,err := sumOfFibonaci(no)
	if err!=nil{
		fmt.Println("Error occured : ",err)
	}else{
		fmt.Println("\nSum of fibonaci series upto ",no," terms is : ",sum)
	}

}