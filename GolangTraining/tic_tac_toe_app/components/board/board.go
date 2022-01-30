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

	var testCells = []*cell.Cell{} //a slice of cell structure pointers

	for i:=0;i<sizeByUser*sizeByUser;i++{
		newCell := cell.New() //new cell is created with noMark
		testCells = append(testCells,newCell)
	}

	var boardTest = &board{
		nCells: testCells, 
		size: sizeByUser,
	}
	return boardTest //pointer to board
}


func (b*board) Display(){
	index:=0
	for i:=0;i<b.size;i++{
		for j:=0;j<b.size;j++{
			icell := b.nCells[index]	//get a structure pointer of the cell at a particular index
			fmt.Print(" | ",icell.GetMark())	//get the mark of the cell at that index
			index++
		}
		fmt.Printf(" |\n")
	}
}

func (b*board) GetBoard() string{
	strBoard := ""
	for i:=0;i<b.size*b.size;i++{
		icell := b.nCells[i]
		strBoard = strBoard + string(icell.GetMark())
	}
	return strBoard
}

func (b*board) IsFull() bool{
	for i:=0;i<b.size*b.size;i++{
		icell := b.nCells[i]
	    if icell.GetMark()==cell.NoMark{
			return false
		}
	}

	return true
}