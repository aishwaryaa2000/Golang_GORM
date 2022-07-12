package service
import (
	
	"battleShip/components/board"
	
)

func Board(rowSizeByUser,colSizeByUser uint8) *board.Board{
	board := board.New(rowSizeByUser,colSizeByUser)
	board.Display()
	board.BoardInit()
	board.Display()
	return board
}
