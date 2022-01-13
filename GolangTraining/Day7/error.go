package main

import (
	// "errors"
	"fmt"
	// "log"
)

type myError struct{
	errmsg string
}

func (err *myError)Error() string{ //method
	return err.errmsg
}

func main() {
		quotient,err := divide(4,0)
		if err!=nil{
			fmt.Println("Error occured : ",err)
			return 
			// log.Fatal(err)
		}
		fmt.Println("Result is  : ",quotient)
		
}

func divide(numerator, denominator int) (int, error) {
	if denominator == 0 {
		// return 0, errors.New("Cannot divide by zero")
		return 0, &myError{errmsg: "Cannot divide by zero"}
	}
	return numerator/denominator,nil
}