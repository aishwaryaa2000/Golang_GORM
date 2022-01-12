package main
import(
	"Chapter5/datafile"
	"fmt"
	"log"
)	

func main(){
	number,err :=  datafile.GetFloatNoFromFile("../data.txt")
	if err!=nil{
		log.Fatal(err)
	}
	totalSum:=0.0
	for _,value := range number{
		totalSum+=value
	}
	fmt.Printf("Average is : %0.2f ", totalSum/3.0)

}
