package main
import "fmt"

func add(num1 int,num2 int) int{
	return num1+num2
}

func sub(num1 int,num2 int) int{
	return num1-num2
}

func mul(num1 int,num2 int) int{
	return num1*num2
}

func div(num1 int,num2 int) int{
	return num1/num2
}
func mathOperations(num1 int,num2 int,funpassed func(int,int) int) int{
	return funpassed(num1,num2)
}

func main(){
	
   sum:= mathOperations(4,5,add)
   diff:= mathOperations(4,5,sub)
   prod:= mathOperations(4,5,mul)
   Quotient:= mathOperations(10,2,div)

   fmt.Println("Sum of 4 and 5 is :",sum)
   fmt.Println("Difference between 4 and 5 is :",diff)
   fmt.Println("Product of 4 and 5 is :",prod)
   fmt.Println("Quotient of 10/2 is :",Quotient)
}