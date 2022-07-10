package board

import(
	"battleShip/components/cell"
	"fmt"
)

type Board struct{
	NCells [][]*cell.Cell //a slice of cell structure pointers
	rowSize uint8
	colSize uint8

}


func New(rowSizeByUser,colSizeByUser uint8) *Board {


	var testCells = [][]*cell.Cell{} //a slice of cell structure pointers
	
	for i:=0;i<int(rowSizeByUser);i++{
		for j:=0;j<int(colSizeByUser);j++{
		newCell := cell.New() //new cell is created with noMark
		testCells[i][j] = newCell
		}
	}

	var boardTest = &Board{
		NCells: testCells, 
		rowSize:rowSizeByUser,
		colSize:colSizeByUser,
	}
	return boardTest //pointer to board
}

func (b*Board)BoardInit(){
	//Initializing board with 5 random ships of size 5
	noOfShips := 5
	rand.Seed(time.Now().UnixNano())



}


func (b*Board) Display(){
	var i,j uint8
	for i=0;i<b.rowSize;i++{
		for j=0;j<b.colSize;j++{
			icell := b.NCells[i][j]	//get a structure pointer of the cell at a particular index
			fmt.Print(" | ",icell.GetMark())	//get the mark of the cell at that index
		}
		fmt.Printf(" |\n")
	}
}





// import (
// 	"crypto/rand"
// 	"math/big"
// )

// func main() {
// 	for i := 0; i < 4; i++ {
// 		n, _ := rand.Int(rand.Reader, big.NewInt(10))
// 		println(n.Int64())
// 	}
// }














// func (b*Board) GetBoard() string{
// 	strBoard := ""
// 	var i uint8
// 	for i=0;i<b.Size*b.Size;i++{
// 		icell := b.NCells[i]
// 		strBoard = strBoard + string(icell.GetMark())
// 	}
// 	return strBoard
// }

// func (b*Board) IsFull() bool{
// 	for i:=0;i<b.Size*b.Size;i++{
// 		icell := b.NCells[i]
// 	    if icell.GetMark()==cell.NoMark{
// 			return false
// 		}
// 	}

// 	return true
// }