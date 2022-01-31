package board

import(
	"tic_tac_toe_app/components/cell"
	"fmt"
)

type Board struct{
	NCells []*cell.Cell //a slice of cell structure pointers
	Size int
}


func New(sizeByUser int) *Board {

	var testCells = []*cell.Cell{} //a slice of cell structure pointers

	for i:=0;i<sizeByUser*sizeByUser;i++{
		newCell := cell.New() //new cell is created with noMark
		testCells = append(testCells,newCell)
	}

	var boardTest = &Board{
		NCells: testCells, 
		Size: sizeByUser,
	}
	return boardTest //pointer to board
}


func (b*Board) Display(){
	index:=0
	for i:=0;i<b.Size;i++{
		for j:=0;j<b.Size;j++{
			icell := b.NCells[index]	//get a structure pointer of the cell at a particular index
			fmt.Print(" | ",icell.GetMark())	//get the mark of the cell at that index
			index++
		}
		fmt.Printf(" |\n")
	}
}

func (b*Board) GetBoard() string{
	strBoard := ""
	for i:=0;i<b.Size*b.Size;i++{
		icell := b.NCells[i]
		strBoard = strBoard + string(icell.GetMark())
	}
	return strBoard
}

// func (b*Board) IsFull() bool{
// 	for i:=0;i<b.Size*b.Size;i++{
// 		icell := b.NCells[i]
// 	    if icell.GetMark()==cell.NoMark{
// 			return false
// 		}
// 	}

// 	return true
// }