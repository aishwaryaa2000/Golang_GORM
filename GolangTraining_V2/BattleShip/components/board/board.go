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
	XcordinateRandom := (random.Intn(int(b.rowSize)))

	YcordinateRandom := (random.Intn(int(b.colSize)))

	fmt.Println("X is : ",XcordinateRandom,"\nY is : ",YcordinateRandom)
	if YcordinateRandom + 4 < int(b.colSize){
		//Ship towards right
		okBool = placeShip(b.NCells,XcordinateRandom,YcordinateRandom,"right",1)
		if okBool{
			//b.Display()
			noOfShips--
			goto newShipPlacement

		}
	}
	 if YcordinateRandom - 4 >= 0{
		//Ship towards left

		okBool = placeShip(b.NCells,XcordinateRandom,YcordinateRandom,"left",-1)
		if okBool{
			//b.Display()

			noOfShips--
			goto newShipPlacement

		}
	}
	 if XcordinateRandom + 4 < int(b.rowSize){
		//Ship downwards

		okBool = placeShip(b.NCells,XcordinateRandom,YcordinateRandom,"down",1)
		if okBool{
			// b.Display()

			noOfShips--
			goto newShipPlacement

		}

	}
	 if XcordinateRandom - 4 >=0{
		//upwards

		okBool = placeShip(b.NCells,XcordinateRandom,YcordinateRandom,"up",-1)
		if okBool{
			// b.Display()

			noOfShips--
			goto newShipPlacement

		}
	}

	}

	
}



func placeShip(b [][]*cell.Cell,XcordinateRandom,YcordinateRandom int,direction string,orientation int) bool{
//here b me change hona chahiye
	var icell *cell.Cell
	if(direction=="right" || direction=="left"){
		for i:=0;i<5;i++{
			if orientation==-1{
			icell = b[XcordinateRandom][YcordinateRandom - i]
			} else{
				icell = b[XcordinateRandom][YcordinateRandom + i]
			}
			mark := icell.GetMark()
			
			if(mark==cell.BattleShip){
					//alreay marked so cannot place ship here
					return false
				}
		}

		for i:=0;i<5;i++{
			if orientation==-1{
			icell = b[XcordinateRandom][YcordinateRandom - i]
			} else{
				icell = b[XcordinateRandom][YcordinateRandom + i]
			}
			 icell.SetMark(cell.BattleShip)
			
		}
			return true
				
	}else {
		for i:=0;i<5;i++{
			if orientation==-1{
			icell = b[XcordinateRandom-i][YcordinateRandom]
			} else{
				icell = b[XcordinateRandom+i][YcordinateRandom]
			}
			mark := icell.GetMark()
			
			if(mark==cell.BattleShip){
					//alreay marked so cannot place ship here
					return false
				}
			}

			for i:=0;i<5;i++{
				if orientation==-1{
				icell = b[XcordinateRandom-i][YcordinateRandom]
				} else{
					icell = b[XcordinateRandom+i][YcordinateRandom]
				}
				icell.SetMark(cell.BattleShip)				
			}

			return true
	}

}


func (b*Board) DisplayHitMiss(){
	var i,j uint8
	for i=0;i<b.rowSize;i++{
		for j=0;j<b.colSize;j++{
			icell := b.NCells[i][j]	//get a structure pointer of the cell at a particular index
			if icell.GetMark()==cell.BattleShip{ 
				fmt.Print(" |    ")
			}else{
				fmt.Print(" | ",icell.GetMark())	//get the mark of the cell at that index
			}
		}
		fmt.Printf(" |\n")
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


