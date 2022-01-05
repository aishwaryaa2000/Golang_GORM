package main
import "fmt"

func main(){
	for i:=0;i<10;i++{
		if i==5{
			continue
		}
		if i==7{
			break
		}
		fmt.Println("Iterating : ",i)
	}

fmt.Println("Outside for loop")
}