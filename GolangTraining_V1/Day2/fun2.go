package main
import "fmt"

func main(){
    userNumber := addOneTo(9)
	fmt.Println(userNumber)
}

func addOneTo(number int) int {
    number++
	return number
}

// func functionName(variableName dataType) returnType {
