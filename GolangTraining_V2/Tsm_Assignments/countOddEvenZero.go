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

/*
If user enters 100 then count of odd=1 and count of zero=2

If user enters 0001 then the program considers input as 1 
so count of odd=1 and count of zero=0
Here the starting zeros 000 are neglected.
This case can be handled by considering the input as a string input i.e "0001"
Refer countOddEvenZero_edgeCases.go program 

If user enters a negative number then that case is handled by making it positive

*/