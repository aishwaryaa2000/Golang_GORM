package main

import (
	"fmt"
	// "reflect"
	"tic_tac_toe_app/service/gameService"
	"bufio"
	"os"
	"strings"
  
)

func main(){
	in := bufio.NewReader(os.Stdin)
	
	begin:
	fmt.Print("Enter size of board : ")
	var size uint8
	//use bufio
	_,err := fmt.Scanln(&size)
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
	fmt.Print("Hey player 1,enter your name : ")
	name1, _ := in.ReadString('\n')
	name1 = strings.TrimRight(name1, "\r\n")

	fmt.Print("Hey player 2,enter your name : ")
	name2, _ := in.ReadString('\n')
	name2 = strings.TrimRight(name2, "\r\n")


	player1,player2 := gameService.Player(name1,name2)
	fmt.Println(player1.Name,",you have been assigned",player1.Mark,"mark")
	fmt.Println(player2.Name,",you have been assigned",player2.Mark,"mark\n")
	
	gameService.GameStart(board,player1,player2)


}