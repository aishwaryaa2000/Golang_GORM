/*
5 4 3 2 1 means ki 15cells toh we want 
3*5 =15 me we can place 5 ships of size 5 4 3 2 1
but 4*4 =16 me we cannot place bcoz either rowSize or ColSize should atleast be 5 
so that we can accomodate a ship with 5 size 

*/

package main

import(
	"battleShip/service"
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main(){
	in := bufio.NewReader(os.Stdin)

	
	begin1 : 
	fmt.Print("Enter number of rows of the board : ")
	var rowSize uint8
	//use bufio
	_,err1 := fmt.Scanln(&rowSize)
	if err1!=nil{
		fmt.Println("Please enter an integer for board")
		goto begin1
	}
	fmt.Print("Enter number of columns of the board : ")
	

	begin2 : 
	var colSize uint8

	fmt.Print("Enter number of columns of the board : ")
	_,err2 := fmt.Scan(&colSize)
	if err2!=nil{
		fmt.Println("Please enter an integer for board")
		goto begin2
	}
	//use bufio
	var minCol uint8 = 15/uint8(rowSize)
	if rowSize<5 && minCol<5{
		/*but 4*4 =16 which is greater than 15 but we cannot place 
		bcoz either rowSize or ColSize should atleast be 5 
		so that we can accomodate a ship with 5 size */
		minCol = 5
	}
	if minCol>colSize{
		fmt.Println("You should have a minimum of ",minCol," to place 5 ships of size 5,4,3,2 and 1")
		goto begin2

	}
	board := service.Board(rowSize,colSize)
	
	fmt.Print("Hey player,enter your name : ")
	name, _ := in.ReadString('\n')
	name = strings.TrimRight(name, "\r\n")

	player := service.Player(name)
	service.GameStart(board,player)


}





