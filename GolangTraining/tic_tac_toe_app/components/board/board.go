package board

import(
	"tic_tac_toe_app/components/cell"
	"fmt"
)

var board = []*cell.Cell{}  //slice of cell structure pointer

func NewBoard(size int) []*cell.Cell{
	for i:=0;i<size*size;i++{
		icell := cell.NewCell()
		board=append(board,icell)
	}
	return board
}

func Display(board []*cell.Cell,size int){
	for i:=0;i<size;i++{
		for j:=0;j<size;j++{
			fmt.Print("|",*board[i])
		}
		fmt.Printf("|\n")
	}
}

func GetCellStatus(board []*cell.Cell,i int) *cell.Cell{
	return board[i]
}