package main
import "fmt"

func main(){
	var userNumber int
	test :
	fmt.Println("Enter a Number")
	fmt.Scanln(&userNumber)
	fmt.Println("The number user has entered is : ",userNumber)
	
	if userNumber >100{
		fmt.Println("Number greater than 100.Please enter a new number")
		goto test
	}else if userNumber<50{
		fmt.Println("Number smaller than 50")
	}else{
		fmt.Println("Number in between 50 and 100.Exiting..")
		return
	}
}


// label :
// goto label