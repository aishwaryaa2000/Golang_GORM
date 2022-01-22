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
	boardn:=board.New(size)
	boardn.Display()

}