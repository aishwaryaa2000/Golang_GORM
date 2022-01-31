package result

import(
	// "tic_tac_toe_app/components/board"
	"math"

)
type Status uint8
const (
	Win Status  = 0
	Tie Status   = 2
	Running Status  = 1
)

func GetStatus(board string) Status{
	if checkColWin(board) || checkRowWin(board) || checkPriDiaWin(board) || checkSecDiaWin(board) {
		return Win
	}else if checkTie(board){
		return Tie
	}
	return Running
}

func checkRowWin(board string) bool{
	strLen := len(board)
	boardSize := int(math.Sqrt(float64(strLen)))
	for i:=0;i<strLen;i+=boardSize{
		ok := checkIfSame(board[i:i+boardSize])
		if ok==true{
			return ok
		}
	}
	return false

}

func checkColWin(board string)bool{
	strLen := len(board)
	boardSize := int(math.Sqrt(float64(strLen)))
	for i:=0;i<boardSize;i++{
		colValue:=""
		for k:=0;k<boardSize;k++{
			colValue += string(board[i+k*boardSize])
		}
		ok := checkIfSame(colValue)
		if ok==true{
			return ok
		}
	}
	return false
}

func checkPriDiaWin(board string)bool{
	strLen := len(board)
	boardSize := int(math.Sqrt(float64(strLen)))
	diagonalValue := ""
	for i:=0;i<boardSize*boardSize;i+=boardSize+1{
		diagonalValue+= string(board[i])
	}

	return checkIfSame(diagonalValue)
}

func checkSecDiaWin(board string)bool{
	strLen := len(board)
	boardSize := int(math.Sqrt(float64(strLen)))
	diagonalValue := ""
	for i:=boardSize-1;i<boardSize*boardSize;i+=boardSize-1{
		diagonalValue+= string(board[i])
	}

	return checkIfSame(diagonalValue)
}

func checkIfSame(strMark string) bool{
	if strMark=="XXX" || strMark=="OOO"{
		return true
	}
	return false
}

func checkTie(board string) bool{
	for i:=0;i<len(board);i++{
		if(board[i]=='-'){
			return false
		}
	}
	return true
}