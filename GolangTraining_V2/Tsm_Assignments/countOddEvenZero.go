package main
import(
	"fmt"
	"log"
)
func count(num int) (uint64,uint64,uint64){

	no:=num
	var digit int 
	var countZero,countEven,countOdd uint64
	if no==0{
		return 0,0,1
	}
	if no<0{
		no=no*-1
		//if the number is negative then make it positive first so that goes into the for loop
	}

	for ;no>0;{
		digit =  no%10
		if digit==0{
			countZero++
		}else if digit%2==0{
			countEven++
		}else{
			countOdd++
		}
		no=no/10
	}
	return countOdd,countEven,countZero

}

func main(){
	var no int 
   
    fmt.Print("Enter the number  : ")
	_, err := fmt.Scanf("%d", &no)
	if err!=nil{
		//If user doesn't enter a number then simply raise an error and exit
		//Example-If user enters an alphabet then raise an error
		log.Fatal(err)
	}	
	odd,even,zero := count(no)
	fmt.Println("Count of odd : ",odd,"\nCount of even : ",even,"\nCount of zero : ",zero)
}