package service

import (
	"battleShip/components/board"
	"battleShip/components/cell"
	"math/rand"
	"fmt"
	"time"
)

func MakeBoard() *board.Board{
	var rowSize uint8
	var colSize uint8
	begin1 : 
	fmt.Print("Enter number of rows of the board : ")
	_,err1 := fmt.Scanln(&rowSize)
	if err1!=nil || rowSize<5{
		fmt.Println("Please enter an integer greater than 4 for board")
		goto begin1
	}
	
	begin2 : 
	fmt.Print("Enter number of columns of the board : ")
	_,err2 := fmt.Scanln(&colSize)
	if err2!=nil || colSize<5{
		fmt.Println("Please enter an integer greater than 4")
		goto begin2
	}
	board := board.New(rowSize,colSize)
	BoardInit(board) //Initializing board with 5 random ships of size 5 4 3 2 1
	board.Display()
	return board
}

func BoardInit(b *board.Board) {
	//Initializing board with 5 random ships of size 5 4 3 2 1 horizontally or vertically
	shipSize := 5

	for shipSize > 0 {
		seed := rand.NewSource(time.Now().UnixNano())
		random := rand.New(seed)
		XcordinateRandom := (random.Intn(int(b.RowSize)))
		YcordinateRandom := (random.Intn(int(b.ColSize)))
		orientation := random.Intn(2)
		//This is to check if vertical placement should be done first then horizontal or vice versa thereby increasing randomness
		if orientation==1{
				rowStart, rowEnd, okBool := checkIfHorizontalOrVertical(b.NCells, XcordinateRandom, YcordinateRandom, int(b.RowSize), shipSize,"vertical")
				if okBool {
					placeShip(b.NCells, rowStart, rowEnd, YcordinateRandom, "vertical")
					shipSize--
					continue
				}
				colStart, colEnd, okBool := checkIfHorizontalOrVertical(b.NCells, XcordinateRandom, YcordinateRandom, int(b.ColSize), shipSize,"horizontal")
				if okBool {
					placeShip(b.NCells, colStart, colEnd, XcordinateRandom, "horizontal")
					shipSize--
					continue
				}
		}else{
				colStart, colEnd, okBool := checkIfHorizontalOrVertical(b.NCells, XcordinateRandom, YcordinateRandom, int(b.ColSize), shipSize,"horizontal")
				if okBool {
					placeShip(b.NCells, colStart, colEnd, XcordinateRandom, "horizontal")
					shipSize--
					continue	
				}
				rowStart, rowEnd, okBool := checkIfHorizontalOrVertical(b.NCells, XcordinateRandom, YcordinateRandom, int(b.RowSize), shipSize,"vertical")
				if okBool {
					placeShip(b.NCells, rowStart, rowEnd, YcordinateRandom, "vertical")
					shipSize--
					continue
				}

  			}
    }

}

func checkIfHorizontalOrVertical(b [][]*cell.Cell, XcordinateRandom, YcordinateRandom, rowOrColSize, shipSize int,orientation string) (int, int, bool){
	var icell *cell.Cell
	icell = b[XcordinateRandom][YcordinateRandom]
	if icell.GetMark()==cell.BattleShip{
		return -1, -1, false
	}

	upOrLeftMin, downOrRightMax := YcordinateRandom, YcordinateRandom

	if orientation=="vertical"{
		upOrLeftMin, downOrRightMax  = XcordinateRandom, XcordinateRandom
	}

	for upOrLeftMin >= 0 {
		if orientation=="vertical"{
			icell = b[upOrLeftMin][YcordinateRandom]
		}else{
			icell = b[XcordinateRandom][upOrLeftMin]
		}
		if icell.GetMark() == cell.BattleShip {
			break
		}
		upOrLeftMin--
	}
	upOrLeftMin++

	for downOrRightMax < rowOrColSize {
		if orientation=="vertical"{
			icell = b[downOrRightMax][YcordinateRandom]
		}else{
			icell = b[XcordinateRandom][downOrRightMax]
		}
		if icell.GetMark() == cell.BattleShip {
			break
		}
		downOrRightMax++
	}
	downOrRightMax--

	if downOrRightMax-upOrLeftMin+1 >= shipSize {
		return upOrLeftMin, upOrLeftMin + shipSize, true
	}

	return -1, -1, false
	
}

func placeShip(b [][]*cell.Cell, min, max, cordinate int, direction string) {

	var icell *cell.Cell
		for i := min; i < max; i++ {
			if direction == "horizontal" {
			    icell = b[cordinate][i]
			}else{
				icell = b[i][cordinate]
			}
			icell.SetMark(cell.BattleShip)
		}

}
