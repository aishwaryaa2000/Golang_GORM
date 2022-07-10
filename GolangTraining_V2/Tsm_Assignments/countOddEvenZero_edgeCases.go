/* 
If user enters 0001 and input is taken as int then the program considers input as 1 
so count of odd=1 and count of zero=0
Here the starting zeros 000 are neglected.

This case can be handled by considering the input as a string input i.e "0001"

Hence if input is 0001
OUTPUT
Count of odd : 1
Count of zero : 3
Count of even : 0

Here,if user enters a negative number then that case is handled by the CONTINUE statement i.e by 
neglecting the - sign
Hence if input is -000234
OUTPUT
Count of odd :  1
Count of even :  2
Count of zero :  3

If user enters any character then error is raised and program exits

*/

package main
import(
	"fmt"
	"log"
	"strconv"
)
func count(num string) (uint64,uint64,uint64){

	var countZero,countEven,countOdd uint64
	for i:=0;i<len(num);i++{
		if i==0 && num[i]=='-'{
			//If user enters a negative number then ignore the - sign
			//Neglect the - sign at index 0 of the string
			continue
		}
		digit,err := strconv.Atoi(string(num[i]))
		if err!=nil{
			//If user has entered a character then raise error and exit
			fmt.Println("Number entered is invalid")
			log.Fatal(err)
		}
		if digit==0{
			countZero++
		}else if digit%2==0{
			countEven++
		}else{
			countOdd++
		}

	}

	return countOdd,countEven,countZero

}

func main(){
	var no string 
   
    fmt.Print("Enter the number  : ")
	fmt.Scanln(&no)

	odd,even,zero := count(no)
	fmt.Println("Count of odd : ",odd,"\nCount of even : ",even,"\nCount of zero : ",zero)
}

