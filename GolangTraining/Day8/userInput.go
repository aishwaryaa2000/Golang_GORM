// bufio and fmt package for taking input
package main

import (
	"bufio"
	"fmt"
	// "log"
	"os"
)

func main(){
	fmt.Printf("Enter something : ")
	reader := bufio.NewReader(os.Stdin)
	input,_ := reader.ReadString('\n') //Read until new line is encountered
	fmt.Println("Input is : ",input)

	fmt.Printf("Enter something : ")
	var number int //default value is 0
	_,err := fmt.Scan(&number) 
	if err!=nil{
		// log.Fatal("Error occurred : ",err)
		fmt.Println("Error has occured : ",err)
	}
	fmt.Println("Input via fmt is : ",number)

}