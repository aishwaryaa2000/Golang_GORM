package main

import (
	// "fmt"
	"fmt"
	// "reflect"
	"tic_tac_toe_app/components/board"
)

func main(){

	begin:
	fmt.Print("Enter size of board : ")
	var size int
	_,err := fmt.Scan(&size)
	if err!=nil{
		fmt.Println("Please enter an integer for board")
		goto begin
	}
	if size<3{
		fmt.Println("Please enter size greater than 2")
		goto begin
	}
	boardn:=board.NewBoard(size)
	board.Display(boardn,size)
	var player1Name, _ string
	fmt.Print("Enter your name : ")
	fmt.Scan(&player1Name)
	//get cell status of board[i] by calling
	//func GetCellStatus(board []*cell.Cell,i int)*cell.Cell  -> board package
	//then mark cell by calling  
	//(c *Cell) markCell(mark string) (*Cell, error) -> cell package


}