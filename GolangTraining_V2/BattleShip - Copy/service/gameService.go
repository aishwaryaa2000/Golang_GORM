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
	board.BoardInit()
	fmt.Println("\n\n\n Final board after placing the ships randomly")
	board.Display()
	return board
}

func Player(name string) (*player.Player){

	PlayerNow:= player.New(name)
	return PlayerNow
}

func GameStart(b *board.Board,currentPlayer *player.Player){
	var X,Y uint8
	shipCells := 25
	noOfTries := 50

	fmt.Println("You only have ",noOfTries," to attack all the ships")
	for ;shipCells>0 && noOfTries>0;{
		fmt.Print(currentPlayer.Name," enter X cordinate where you wish to attack : ")
		fmt.Scan(&X)
		fmt.Print(currentPlayer.Name," enter Y cordinate where you wish to attack : ")
		fmt.Scan(&Y)

		icell := b.NCells[X][Y]
		if icell.GetMark()==cell.BattleShip{
			icell.SetMark(cell.Hit)
			fmt.Println("Hurray! You hit the ship")
			b.DisplayHitMiss()
			shipCells--
			noOfTries--

		}else if icell.GetMark()==cell.NoMark{
			icell.SetMark(cell.Miss)
			fmt.Println("Oh no,you missed the ship")
			b.DisplayHitMiss()
			noOfTries--
		}else{
			fmt.Println("Already attacked this ship")
		}

	}

	fmt.Println("GAME OVER!\n\n")
	if shipCells==0{
		fmt.Println("You won the game!\nYou needed ",noOfTries," of tries to win the game")
	}else{
		fmt.Println("Oh no!You have exhausted all your tries")
	}
	
}
