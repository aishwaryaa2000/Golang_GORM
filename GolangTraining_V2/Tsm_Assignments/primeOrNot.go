//Time complexity is Sqrt(N)
package main
import(
	"fmt"
	"math"
	"log"

)
func primeOrNot(num int) bool{

	if num<2{
		//prime numbers are always greater than 1
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
        if (num % i == 0) {
			//There exists a divisor other than 1 so not a prime number
			return false
        }
    }

	return true

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

	if primeOrNot(no){
		fmt.Println("It is a prime number")
	}else{
		fmt.Println("It is not a prime number")
	}
}