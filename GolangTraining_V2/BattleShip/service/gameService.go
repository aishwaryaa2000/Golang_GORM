package service

import (
	"battleShip/components/board"
	"fmt"
)

func Board(rowSizeByUser,colSizeByUser uint8) *board.Board{
	board := board.New(rowSizeByUser,colSizeByUser)
	board.Display()
	board.BoardInit()

	fmt.Println("\n\n\n final")
	board.Display()
	return board
}
