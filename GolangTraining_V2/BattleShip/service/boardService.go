package service

import (
	"battleShip/components/board"
	"battleShip/components/cell"
	"fmt"
	"math/rand"
	"time"
)

func BoardInit(b *board.Board) {
	//Initializing board with 5 random ships of size 5
	shipSize := 5
	newShipPlacement:
	for shipSize > 0 {
		seed := rand.NewSource(time.Now().UnixNano())
		random := rand.New(seed)
		XcordinateRandom := (random.Intn(int(b.RowSize)))
		YcordinateRandom := (random.Intn(int(b.ColSize)))

		fmt.Println("X is : ", XcordinateRandom, "\nY is : ", YcordinateRandom)
		orientation := random.Intn(2)
if orientation==1{
	fmt.Println("vertical phele")

	rowMin, rowMax, okBool := checkIfVertical(b.NCells, XcordinateRandom, YcordinateRandom, int(b.RowSize), shipSize)
	if okBool {

		placeShip(b.NCells, rowMin, rowMax, YcordinateRandom, "vertical")
		fmt.Println("Ship size placed ", shipSize)
		b.Display()
		shipSize--
		goto newShipPlacement
	}

	colMin, colMax, okBool := checkIfHorizontal(b.NCells, XcordinateRandom, YcordinateRandom, int(b.ColSize), shipSize)
	if okBool {

		placeShip(b.NCells, colMin, colMax, XcordinateRandom, "horizontal")
		fmt.Println("Ship size placed ", shipSize)
		b.Display()
		shipSize--
		goto newShipPlacement
	}
}else{
	fmt.Println("horizontal phele")
	colMin, colMax, okBool := checkIfHorizontal(b.NCells, XcordinateRandom, YcordinateRandom, int(b.ColSize), shipSize)
	if okBool {

		placeShip(b.NCells, colMin, colMax, XcordinateRandom, "horizontal")
		fmt.Println("Ship size placed ", shipSize)
		b.Display()
		shipSize--
		goto newShipPlacement	
	}

	rowMin, rowMax, okBool := checkIfVertical(b.NCells, XcordinateRandom, YcordinateRandom, int(b.RowSize), shipSize)
	if okBool {

		placeShip(b.NCells, rowMin, rowMax, YcordinateRandom, "vertical")
		fmt.Println("Ship size placed ", shipSize)
		b.Display()
		shipSize--
		goto newShipPlacement
	}

  }

 }

}

func checkIfVertical(b [][]*cell.Cell, XcordinateRandom, YcordinateRandom, rowSize, shipSize int) (int, int, bool) {
	//row index dhekna hoga
	var icell *cell.Cell
	var mark cell.Mark

	icell = b[XcordinateRandom][YcordinateRandom]
	if icell.GetMark()==cell.BattleShip{
		return -1, -1, false
	}

	var (
		upMin, downMax int = XcordinateRandom, XcordinateRandom
	)

	for upMin >= 0 {
		icell = b[upMin][YcordinateRandom]
		mark = icell.GetMark()
		if mark == cell.BattleShip {
			break
		}
		upMin--
	}

	upMin++
	for downMax < rowSize {
		icell = b[downMax][YcordinateRandom]
		mark = icell.GetMark()
		if mark == cell.BattleShip {
			break
		}
		downMax++
	}

	downMax--

	if downMax-upMin+1 >= shipSize {
		//rowMin rowMax bool
		fmt.Println("Vertical:", upMin, downMax)
		return upMin, upMin + shipSize, true
	} else {
		return -1, -1, false
	}

}

func checkIfHorizontal(b [][]*cell.Cell, XcordinateRandom, YcordinateRandom, colSize, shipSize int) (int, int, bool) {
	var icell *cell.Cell
	var mark cell.Mark

	icell = b[XcordinateRandom][YcordinateRandom]
	if icell.GetMark()==cell.BattleShip{
		return -1, -1, false
	}

	var (
		leftMin, rightMax int = YcordinateRandom, YcordinateRandom
	)
	for leftMin >= 0 {
		icell = b[XcordinateRandom][leftMin]
		mark = icell.GetMark()
		if mark == cell.BattleShip {
			break
		}
		leftMin--
	}

	leftMin++
	for rightMax < colSize {
		icell = b[XcordinateRandom][rightMax]
		mark = icell.GetMark()
		if mark == cell.BattleShip {
			break
		}
		rightMax++
	}

	rightMax--

	if rightMax-leftMin+1 >= shipSize {
		return leftMin, leftMin + shipSize, true
	} 
	
	return -1, -1, false
	
}

func placeShip(b [][]*cell.Cell, min, max, cordinate int, direction string) {
	//here b me change hona chahiye
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
