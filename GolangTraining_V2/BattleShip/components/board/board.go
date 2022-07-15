package board

import (
	"battleShip/components/cell"
	"fmt"
)

type Board struct {
	NCells  [][]*cell.Cell //a slice of cell structure pointers
	rowSize uint8
	colSize uint8
}

func New(rowSizeByUser, colSizeByUser uint8) *Board {

	var testCells = [][]*cell.Cell{} //a slice of cell structure pointers

	for i := 0; i < int(rowSizeByUser); i++ {
		ithRow := []*cell.Cell{}
		for j := 0; j < int(colSizeByUser); j++ {
			newCell := cell.New() //new cell is created with noMark
			ithRow = append(ithRow, newCell)
		}
		testCells = append(testCells, ithRow)

	}

	var boardTest = &Board{
		NCells:  testCells,
		rowSize: rowSizeByUser,
		colSize: colSizeByUser,
	}
	return boardTest //pointer to board
}
func (b *Board) GetRowColSize()(uint8,uint8){
	return b.rowSize,b.colSize
}
func (b *Board) DisplayHitMiss() {
	var i, j uint8
	for i = 0; i < b.rowSize; i++ {
		for j = 0; j < b.colSize; j++ {
			icell := b.NCells[i][j] //get a structure pointer of the cell at a particular index
			if icell.GetMark() == cell.BattleShip {
				fmt.Print(" |    ")
			} else {
				fmt.Print(" | ", icell.GetMark()) //get the mark of the cell at that index
			}
		}
		fmt.Printf(" |\n")
	}
}

func (b *Board) Display() {
	var i, j uint8
	for i = 0; i < b.rowSize; i++ {
		for j = 0; j < b.colSize; j++ {
			icell := b.NCells[i][j]           //get a structure pointer of the cell at a particular index
			fmt.Print(" | ", icell.GetMark()) //get the mark of the cell at that index
		}
		fmt.Printf(" |\n")
	}
}

