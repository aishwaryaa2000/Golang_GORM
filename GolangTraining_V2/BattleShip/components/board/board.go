package board

import(
	"battleShip/components/cell"
	"fmt"
	"math/rand"
	"time"
)

type Board struct{
	NCells [][]*cell.Cell //a slice of cell structure pointers
	rowSize uint8
	colSize uint8

}


func New(rowSizeByUser,colSizeByUser uint8) *Board {


	var testCells = [][]*cell.Cell{} //a slice of cell structure pointers
	
	for i:=0;i<int(rowSizeByUser);i++{
		ithRow := []*cell.Cell{}
		for j:=0;j<int(colSizeByUser);j++{
			newCell := cell.New() //new cell is created with noMark
			ithRow = append(ithRow, newCell)
		}
		testCells = append(testCells, ithRow)

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
	newShipPlacement:
	okBool := false
	for ;noOfShips>0 ;{
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	XcordinateRandom := uint8(random.Intn(int(b.rowSize)))

	YcordinateRandom := uint8(random.Intn(int(b.colSize)))

	fmt.Println("X is : ",XcordinateRandom,"\nY is : ",YcordinateRandom)
	if YcordinateRandom + 4 < b.colSize{
		//Ship towards right
		fmt.Println("inside right")
		okBool = placeShip(b.NCells,XcordinateRandom,YcordinateRandom,"right",1)
		if okBool{
			b.Display()
			noOfShips--
			goto newShipPlacement

		}
	}
	 if YcordinateRandom - 4 >= 0{
		//Ship towards left
		fmt.Println("inside left")

		okBool = placeShip(b.NCells,XcordinateRandom,YcordinateRandom,"left",-1)
		if okBool{
			b.Display()

			noOfShips--
			goto newShipPlacement

		}
	}
	 if XcordinateRandom + 4 < b.rowSize{
		//Ship downwards
		fmt.Println("inside downwards")

		okBool = placeShip(b.NCells,XcordinateRandom,YcordinateRandom,"down",1)
		if okBool{
			b.Display()

			noOfShips--
			goto newShipPlacement

		}

	}
	 if XcordinateRandom - 4 >=0{
		//upwards
		fmt.Println("inside upwards")

		okBool = placeShip(b.NCells,XcordinateRandom,YcordinateRandom,"up",-1)
		if okBool{
			b.Display()

			noOfShips--
			goto newShipPlacement

		}
	}

	}

	
}



func placeShip(b [][]*cell.Cell,XcordinateRandom,YcordinateRandom uint8,direction string,orientation int) bool{
//here b me change hona chahiye
	var icell *cell.Cell
	var i uint8
	if(direction=="right" || direction=="left"){
		for i=0;i<5;i++{
			if orientation==-1{
			icell = b[XcordinateRandom][YcordinateRandom - i]
			} else{
				icell = b[XcordinateRandom][YcordinateRandom + i]
			}
			errorSetMark := icell.SetMark(cell.BattleShip)
				if errorSetMark != nil{
					return false
				}
			}

			return true
				
	}else {
		for i=0;i<5;i++{
			if orientation==-1{
			icell = b[XcordinateRandom-i][YcordinateRandom]
			} else{
				icell = b[XcordinateRandom+i][YcordinateRandom]
			}
			errorSetMark := icell.SetMark(cell.BattleShip)
				if errorSetMark != nil{
					return false
				}
			}

			return true
	}

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
		// n, _ := rand.Int(rand.Reader, big.NewInt(10))
		// println(n.Int64())
// 	}
// }













/* 
var testCells [][]*string //a slice of cell structure pointers
	fmt.Println(reflect.TypeOf(testCells))
	str := "aishu"
	p := &str
	testCells[0][0] = p

	var testCells []*string //a slice of cell structure pointers
	fmt.Println(reflect.TypeOf(testCells))
	str := "aishu"
	p := &str
	testCells[0] = p


*/
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