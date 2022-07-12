/* 
Using maps

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
	numberSlice := make([]int,0,10)
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
		numberSlice = append(numberSlice,userNumber)
	}

	getCount(numberSlice)
}

func getCount(number []int) {

	var numberCountMap = make(map[int]int)
	for _, value := range number {
		numberCountMap[value]++
	}

	for num, count := range numberCountMap {
		if count==1{
        fmt.Println( num, " occured ",count," time")
		//1 time
		}else{
			fmt.Println( num, " occured ",count," times")
		//times
		}
    }
	
}