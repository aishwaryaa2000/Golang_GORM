package main

import (
	// "fmt"
	"fmt"
	// "reflect"
	// "tic_tac_toe_app/components/board"
	"tic_tac_toe_app/service/gameService"


)

func main(){

	begin:
	fmt.Print("Enter size of board : ")
	var size int
	//use bufio
	_,err := fmt.Scan(&size)
	if err!=nil{
		fmt.Println("Please enter an integer for board")
		goto begin
	}
	if size<3{
		fmt.Println("Please enter size greater than 2")
		goto begin
	}
	board := gameService.Board(size)
	// boardn.Display()
	var name1,name2 string
	fmt.Print("Hey player 1,enter your name : ")
	fmt.Scan(&name1)
	fmt.Print("Hey player 2,enter your name : ")
	fmt.Scan(&name2)
	player1,player2 := gameService.Player(name1,name2)
	fmt.Println(player1.Name,",you have been assigned",player1.Mark,"mark")
	fmt.Println(player2.Name,",you have been assigned",player2.Mark,"mark\n")
	gameService.GameStart(board,player1,player2)


}