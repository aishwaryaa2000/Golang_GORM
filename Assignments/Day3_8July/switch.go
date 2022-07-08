package main

import (
	"fmt"
	"log"
)

func main(){
	var ptr int
	fmt.Print("Enter your pointer : ") //assume ptr to be <=10
    _, err := fmt.Scanf("%d", &ptr)
	if err!=nil{
		log.Fatal(err)
	}

	switch ptr {
	case 10:
		fmt.Println("Grade is : A+") 
	case 9,8:
		fmt.Println("Grade is : A")
	case 7:
		fmt.Println("Grade is : B")
	case 6:
		fmt.Println("Grade is : C")
	default:
		fmt.Println("Grade is : D")

	}
}