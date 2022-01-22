package board

import(
	"tic_tac_toe_app/components/cell"
	"fmt"
)

type board struct{
	nCells []*cell.Cell //a slice of cell structure pointers
	size int
}


func New(sizeByUser int) *board {

	var testCells = []*cell.Cell{} 

	for i:=0;i<sizeByUser*sizeByUser;i++{
		newCell := cell.New() //new cell is created with noMark
		testCells = append(testCells,newCell)
	}

	var boardTest = &board{
		nCells: testCells, 
		size: sizeByUser,
	}
	return boardTest //pointer to player
}


func (b*board) Display(){
	index:=0
	for i:=0;i<b.size;i++{
		for j:=0;j<b.size;j++{
			icell := b.nCells[index]
			fmt.Print(" | ",icell.GetMark())
			index++
		}
		fmt.Printf(" |\n")
	}
}
