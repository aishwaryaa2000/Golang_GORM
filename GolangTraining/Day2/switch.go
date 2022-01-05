package main

import "fmt"

func main(){
	cgpa:=8
	switch cgpa {
	case 10:
		fmt.Println("Excellent") 
	case 9,8:
		fmt.Println("Very good")
	case 7:
		fmt.Println("Good")
	case 6:
		fmt.Println("Average")
	default:
		fmt.Println("Bad")

	}
}

// if case 10 is true then rest cases won't be evaluated
// and execution of the program will directly go outside switch block so no BREAK statement required

// for multiple cases and same body, use comma 