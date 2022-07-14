package service

import (
	"battleShip/components/board"
	"battleShip/components/player"
	"battleShip/components/cell"

	"fmt"
)

func Board(rowSizeByUser,colSizeByUser uint8) *board.Board{
	board := board.New(rowSizeByUser,colSizeByUser)
	board.Display()
	BoardInit(board)
	fmt.Println("\n\n\n Final board after placing the ships randomly")
	board.Display()
	return board
}

func Player(name string) (*player.Player){

	PlayerNow:= player.New(name)

	return PlayerNow
}

func takeInput(name string,rowSize,colSize uint8 )(uint8,uint8){
	var X,Y uint8

	begin1:
		fmt.Print(name," enter X cordinate where you wish to attack : ")
		_,err1 := fmt.Scan(&X)
		if err1!=nil{
			fmt.Println("Please enter an integer for board lesser than ",rowSize)
			goto begin1
		}
		if X>= rowSize {
			fmt.Println("Please enter an integer lesser than ",rowSize)
			goto begin1
		}
	
		fmt.Print(name," enter Y cordinate where you wish to attack : ")
		_,err2 := fmt.Scan(&Y)
		if err2!=nil{
			fmt.Println("Please enter an integer for board lesser than ",colSize)
			goto begin1
		}
		
		if Y>= colSize {
			fmt.Println("Please enter an integer lesser than ",colSize)
			goto begin1
		}

	return X,Y
}

func GameStart(b *board.Board,currentPlayer *player.Player){
	var X,Y uint8
	shipCells := 15
	noOfTries := 50


	fmt.Println("You only have ",noOfTries," to attack all the ships")
	for ;shipCells>0 && noOfTries>0;{
		X,Y = takeInput(currentPlayer.Name,b.RowSize,b.ColSize)
		icell := b.NCells[X][Y]
		if icell.GetMark()==cell.BattleShip{
			icell.SetMark(cell.Hit)
			fmt.Println("Hurray! You hit the ship")
			b.DisplayHitMiss()
			currentPlayer.Hit++
			shipCells--
			noOfTries--

		}else if icell.GetMark()==cell.NoMark{
			icell.SetMark(cell.Miss)
			fmt.Println("Oh no,you missed the ship")
			b.DisplayHitMiss()
			currentPlayer.Miss++
			noOfTries--
		}else{
			fmt.Println("Already attacked this ship")
		}

		fmt.Println("Number of hits : ",currentPlayer.Hit," \nNumber of miss : ",currentPlayer.Miss,"\nTrials left : ",noOfTries)
	}

	//Result analysis
	fmt.Println("GAME OVER!\n\n")
	if shipCells==0{
		fmt.Println("You won the game!\nYou needed ",50 - noOfTries," of tries to win the game")
	}else{
		fmt.Println("Oh no!You have exhausted all your tries")
	}
	
}
