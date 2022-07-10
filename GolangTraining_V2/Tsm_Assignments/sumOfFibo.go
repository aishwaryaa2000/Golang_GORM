package main

import(
	"fmt"
	"log"
)

func sumOfFibonaci(noOfTerms int) uint64 {
	//0 1 1 2 3 8...
	if noOfTerms==1{
		return 0
	}else if noOfTerms==2{
		return 1
	}else{
		var(
			sumTotal, firstNum, secondNum,thirdNum uint64= 1,0,1,0
		  )

		for i:=0;i<noOfTerms-2;i++{
			thirdNum = firstNum+secondNum
			sumTotal += thirdNum
			firstNum=secondNum
			secondNum=thirdNum
		}

		return sumTotal

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
	fmt.Println("Sum of fibonaci series upto ",no," terms is : ",sumOfFibonaci(no))

}