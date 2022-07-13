/* 
	Insertion sort 
	O(n^2)

	Merge sort can be used as time complexity is O(n logn) but it uses stack 
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

	sorting(numberSlice)
	fmt.Println("\nAfter sorting the array : ",numberSlice)
	//Slice is always call by ref so changes are reflected

}

func sorting(number []int){
	var j int
	var temp int
    for i := 1;i< len(number); i++ {
        temp = number[i]
        j = i - 1

        for ; j >= 0 && number[j] > temp ;{
            number[j + 1] = number[j]
            j--
        }
        number[j + 1] = temp

    }

}
