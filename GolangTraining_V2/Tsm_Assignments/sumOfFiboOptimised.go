package main
import(
	"math"
	"fmt"
	"log"
	"errors"
)

func fiboNthTerm(noOfTerms int) (uint64,error) {
	//Taking uint64 as return type because we know that the sum will always be positive thereby increasing the range

	if noOfTerms<1{
		//Here we will check noOfTerms<1 since while function call we are passing noOfTerms+1
		err := errors.New("Number of terms cannot be less than zero")
		return 0,err
	}else{
		phi  := (1 + math.Sqrt(5)) / 2
		return uint64(math.Round(math.Pow(phi, float64(noOfTerms)) / math.Sqrt(5))),nil
	}
}

// Driver Code
func main (){
	var no int 
   
    fmt.Print("Enter number of terms : ")
	_, err := fmt.Scanf("%d", &no)
	if err!=nil{
		//If user doesn't enter a number then simply raise an error and exit
		//Example-If user enters an alphabet then raise an 
		log.Fatal(err)
	}
	sum,err := fiboNthTerm(no+1)
	//S(n) = F(n+1) – 1 
	//when we consider fibonacci series as 0 1 1 3 4 where 0 is the 1st term, 1 is the 2nd term and so on...
	
	if err!=nil{
		fmt.Println("Error occured : ",err)
	}else{
		fmt.Println("Sum of fibonaci series upto ",no," terms is : ",sum-1)
	}
}
