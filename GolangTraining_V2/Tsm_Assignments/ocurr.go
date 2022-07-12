/* 
if we store the numbers given by the user directly into the map and update the frequencies simultaneously 
then number of loops required is lesser than occurrence.go


If user enters a negative size of character then error is raised and program exits
If user enters a character instead of a number while entering elements into the slice then error is raised and program exits

If you do not want program to exit whenever error is raised, then GOTO can be used so that the user can re-enter a valid number

*/

package main

import (
	"fmt"
	"log"
)

func main() {
	var numberCountMap = make(map[int]int)
    
    var len,i uint64
	var userNumber int
	fmt.Print("Enter the count of numbers you wish to enter : ")
	_,err := fmt.Scanln(&len)

	if err!=nil{
		//If user enters a character i.e not a number and when user enters a negative number as len is uint
		fmt.Println("Either you entered a negative number or a character")
		log.Fatal(err)
	}

    for i=0;i<len;i++{
		fmt.Print("Enter number",i+1,": ")

		_,err2:=fmt.Scanln(&userNumber)
		if err2!=nil{
			//If user enters a character i.e not a number
			log.Fatal(err2)
		}

		//Storing the numbers into the map and simultaneously storing the count of numbers
		numberCountMap[userNumber]++
		
	}

	displayCount(numberCountMap)

}

func displayCount(NumFreq map[int]int){
	//Displaying the count and numbers stored in the map
	for num, freq := range NumFreq {
		if freq==1{
        fmt.Println( num, " occured ",freq," time")
		//1 time
		}else{
			fmt.Println( num, " occured ",freq," times")
		//times
		}
    }
}