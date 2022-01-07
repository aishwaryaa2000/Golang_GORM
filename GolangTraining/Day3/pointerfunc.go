package main
import "fmt"

var num int 

func main(){
	num=20
	fmt.Println("Address of num in main : ",&num)
	addTwo(num)
	addTwoModified(&num)
	fmt.Println("Value of num in main after calling addTwoModified : ",num)


}

func addTwo(num int)  {
	fmt.Println("Address of num in addTwo : ",&num)
	fmt.Println("Value of num in addTwo : ",num)

}

func addTwoModified(num *int){
	fmt.Println("Address of num in addTwoModified : ",&num)
	fmt.Println("Value stored in num : ",*num)
	fmt.Println("Value stored in num : ",num)

	*num = *num +2
}