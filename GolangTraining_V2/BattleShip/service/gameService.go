package service

import (
	"battleShip/components/board"
	"battleShip/components/player"
	"battleShip/components/cell"
	"fmt"
)

func GameStart(b *board.Board,currentPlayer *player.Player){
	var X,Y uint8
	var (shipCells ,noOfTries uint= 15,50)

	fmt.Println("You only have ",noOfTries," to attack all the ships")
	for ;shipCells>0 && noOfTries>0;{
		X,Y = inputFromPlayerToAttack(currentPlayer.Name,b.RowSize,b.ColSize)
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

	resultAnalysis(shipCells,noOfTries)
	
}

func resultAnalysis(shipCells,noOfTries uint){
	fmt.Println("GAME OVER!\n\n")
	if shipCells==0{
		fmt.Println("You won the game!\nYou needed ",50 - noOfTries," of tries to win the game")
	}else{
		fmt.Println("Oh no!You have exhausted all your tries")
	}
}