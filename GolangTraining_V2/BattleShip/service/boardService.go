package service

import (
	"battleShip/components/board"
	"battleShip/components/cell"
	"math/rand"
	"time"
)

func MakeBoard() *board.Board{

	rowSize,colSize := inputSizeFromPlayer()
	board := board.New(rowSize,colSize)
	BoardInit(board) //Initializing board with 5 random ships of size 5 4 3 2 1
	board.Display()
	return board
}

func randomNumberGenerator(rowSize,colSize uint8)(int,int,int){
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	XcordinateRandom := (random.Intn(int(rowSize)))
	YcordinateRandom := (random.Intn(int(colSize)))
	orientation := random.Intn(2)

	return XcordinateRandom,YcordinateRandom,orientation
}

func BoardInit(b *board.Board) {
	//Initializing board with 5 random ships of size 5 4 3 2 1 horizontally or vertically
	noOfShip := 5	
	rowSize,colSize := b.GetRowColSize()
	
	for noOfShip > 0 {
		XcordinateRandom,YcordinateRandom,orientation := randomNumberGenerator(rowSize,colSize)
		/*Orientation is to check if vertical placement should be done first then horizontal or 
		vice versa thereby increasing randomness*/
		if orientation==1{
				rowStart, rowEnd, okBool := checkIfHorizontalOrVertical(b.NCells, XcordinateRandom, YcordinateRandom, int(rowSize), noOfShip,"vertical")
				if okBool {
					placeShip(b.NCells, rowStart, rowEnd, YcordinateRandom, "vertical")
					noOfShip--
					continue
				}
				colStart, colEnd, okBool := checkIfHorizontalOrVertical(b.NCells, XcordinateRandom, YcordinateRandom, int(colSize), noOfShip,"horizontal")
				if okBool {
					placeShip(b.NCells, colStart, colEnd, XcordinateRandom, "horizontal")
					noOfShip--
					continue
				}
		}else{
				colStart, colEnd, okBool := checkIfHorizontalOrVertical(b.NCells, XcordinateRandom, YcordinateRandom, int(colSize), noOfShip,"horizontal")
				if okBool {
					placeShip(b.NCells, colStart, colEnd, XcordinateRandom, "horizontal")
					noOfShip--
					continue	
				}
				rowStart, rowEnd, okBool := checkIfHorizontalOrVertical(b.NCells, XcordinateRandom, YcordinateRandom, int(rowSize), noOfShip,"vertical")
				if okBool {
					placeShip(b.NCells, rowStart, rowEnd, YcordinateRandom, "vertical")
					noOfShip--
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
