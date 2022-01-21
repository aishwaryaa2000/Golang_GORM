package board

import(
	"tic_tac_toe_app/components/cell"
	"fmt"
)

var board = [10]*cell.Cell{} //an array of 9 structure pointers


func NewBoard() [10]*cell.Cell{
	for i:=0;i<10;i++{
		board[i]= cell.NewCell()
		}
	return board
}

func Display(board [10]*cell.Cell){
	fmt.Println(*board[7]," | ",*board[8]," | ",*board[9])
    fmt.Println("-----|-------|-----")
    fmt.Println(*board[4]," | ",*board[5]," | ",*board[6])
    fmt.Println("-----|-------|-----")
    fmt.Println(*board[1]," | ",*board[2]," | ",*board[3])

}

func GetCellStatus(board [10]*cell.Cell,i int) *cell.Cell{
	return board[i]
}