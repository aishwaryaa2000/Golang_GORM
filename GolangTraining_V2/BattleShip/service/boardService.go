package service

import (
	"battleShip/components/board"
	"battleShip/components/cell"
	"math/rand"
	"time"
)

func MakeBoard() *board.Board {

	rowSize, colSize := inputSizeFromPlayer()
	board := board.New(rowSize, colSize)
	BoardInit(board) //Initializing board with 5 random ships of size 5 4 3 2 1
	board.Display()
	return board
}

func randomNumberGenerator(rowSize, colSize uint8) (int, int, int) {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	XcordinateRandom := (random.Intn(int(rowSize)))
	YcordinateRandom := (random.Intn(int(colSize)))
	orientation := random.Intn(2)

	return XcordinateRandom, YcordinateRandom, orientation
}

func BoardInit(currentBoard *board.Board) {
	//Initializing board with 5 random ships of size 5 4 3 2 1 horizontally or vertically
	noOfShip := 5
	rowSize, colSize := currentBoard.GetRowColSize()

	for noOfShip > 0 {
		XcordinateRandom, YcordinateRandom, orientation := randomNumberGenerator(rowSize, colSize)
		/*Orientation is to check if vertical placement should be done first then horizontal or
		vice versa thereby increasing randomness*/
		if orientation == 1 {
			rowStart, rowEnd, okBool := checkIfHorizontalOrVertical(currentBoard.NCells, XcordinateRandom, YcordinateRandom, int(rowSize), noOfShip, 'V')
			if okBool {
				placeShip(currentBoard.NCells, rowStart, rowEnd, YcordinateRandom, 'V')
				noOfShip--
				continue
			}
			colStart, colEnd, okBool := checkIfHorizontalOrVertical(currentBoard.NCells, XcordinateRandom, YcordinateRandom, int(colSize), noOfShip, 'H')
			if okBool {
				placeShip(currentBoard.NCells, colStart, colEnd, XcordinateRandom, 'H')
				noOfShip--
				continue
			}
		} else {
			colStart, colEnd, okBool := checkIfHorizontalOrVertical(currentBoard.NCells, XcordinateRandom, YcordinateRandom, int(colSize), noOfShip, 'H')
			if okBool {
				placeShip(currentBoard.NCells, colStart, colEnd, XcordinateRandom, 'H')
				noOfShip--
				continue
			}
			rowStart, rowEnd, okBool := checkIfHorizontalOrVertical(currentBoard.NCells, XcordinateRandom, YcordinateRandom, int(rowSize), noOfShip, 'V')
			if okBool {
				placeShip(currentBoard.NCells, rowStart, rowEnd, YcordinateRandom, 'V')
				noOfShip--
				continue
			}

		}
	}

}

func checkIfHorizontalOrVertical(b [][]*cell.Cell, XcordinateRandom, YcordinateRandom, rowOrColSize, shipSize int, orientation rune) (int, int, bool) {
	var icell *cell.Cell
	icell = b[XcordinateRandom][YcordinateRandom]
	if icell.Cell() == cell.BattleShip {
		return -1, -1, false
	}

	upOrLeftMin, downOrRightMax := YcordinateRandom, YcordinateRandom

	if orientation == 'V' {
		upOrLeftMin, downOrRightMax = XcordinateRandom, XcordinateRandom
	}

	for upOrLeftMin >= 0 {
		if orientation == 'V' {
			icell = b[upOrLeftMin][YcordinateRandom]
		} else {
			icell = b[XcordinateRandom][upOrLeftMin]
		}
		if icell.Cell() == cell.BattleShip {
			break
		}
		upOrLeftMin--
	}
	upOrLeftMin++

	for downOrRightMax < rowOrColSize {
		if orientation == 'V' {
			icell = b[downOrRightMax][YcordinateRandom]
		} else {
			icell = b[XcordinateRandom][downOrRightMax]
		}
		if icell.Cell() == cell.BattleShip {
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


