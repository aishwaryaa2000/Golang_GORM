package result

import (
	"math"
	"tic_tac_toe_app/components/cell"


)
type Status uint8

var colWin[] int
var rowWin[] int
var priDiaWin int
var secDiaWin int


func MakeWinArray(boardsize int){
	colWin = make([]int, boardsize)
	rowWin = make([]int, boardsize)

}

const (
	Win Status  = 0
	Tie Status   = 2
	Running Status  = 1
)

func GetStatus(boardSize int ,mark cell.Mark,index,noOfTurns int) Status{

	changeWinArray(index,boardSize,mark)

	if checkColRowWin(boardSize,index) || checkPriSecDiaWin(boardSize,index){
		return Win
	}else if noOfTurns==boardSize*boardSize{
		return Tie
	}
	return Running
}


func changeWinArray(index int,boardsize int,mark cell.Mark){
	virtualColumn := index%boardsize
	virtualRow := index/boardsize
	
	if mark == cell.XMark{ 
		rowWin[virtualRow]++
		colWin[virtualColumn]++
	}else{
		rowWin[virtualRow]--
		colWin[virtualColumn]--
	}

	if virtualColumn+virtualRow==boardsize-1{	
			if mark == cell.XMark{ 
				secDiaWin++
			}else{
				secDiaWin--
			}
	}

	if virtualColumn==virtualRow{	
		if mark == cell.XMark{ 
			priDiaWin++
		}else{
			priDiaWin--
		}
	}
}

func checkColRowWin(boardSize ,index int) bool{
	
	virtualColumn := index%boardSize
	virtualRow := index/boardSize

	if boardSize==int(math.Abs(float64(rowWin[virtualRow]))) ||boardSize==int(math.Abs(float64(colWin[virtualColumn]))){
		return true
	}

	return false
}

func checkPriSecDiaWin(boardSize ,index int)bool{
	
	if int(math.Abs(float64(priDiaWin)))==boardSize || int(math.Abs(float64(secDiaWin)))==boardSize{
		return true
	}
	return false

}

