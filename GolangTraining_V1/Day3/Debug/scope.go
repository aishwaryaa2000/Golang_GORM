package main
import "fmt"

var num int 

func main(){
    num=40
	display()
	fmt.Println("Inside main method",num)
}

func display(){
	var num1 int
	num1=30
	fmt.Println(num1)
	fmt.Println("Inside display function ",num)
}