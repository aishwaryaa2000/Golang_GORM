package main

import(
	"battleShip/service"
)

func main(){

	board := service.MakeBoard()
	player := service.MakePlayer()
	service.GameStart(board,player)

}





















/*
5 4 3 2 1 means ki 15cells toh we want 
3*5 =15 me we can place 5 ships of size 5 4 3 2 1
but 4*4 =16 me we cannot place bcoz either rowSize or ColSize should atleast be 5 
so that we can accomodate a ship with 5 size 


use bufio
	var minCol uint8 = 15/uint8(rowSize)
	if rowSize<5 && minCol<5{
		
		but 4*4 =16 which is greater than 15 but we cannot place 
		bcoz either rowSize or ColSize should atleast be 5 
		so that we can accomodate a ship with 5 size 

		minCol = 5
	}
	if minCol>colSize{
		fmt.Println("You should have a minimum of ",minCol," to place 5 ships of size 5,4,3,2 and 1")
		goto begin2

	}

*/

