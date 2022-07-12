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
	for ;shipCells>0;{
		fmt.Print(currentPlayer.Name," Enter X cordinate nuofmber where you wish to mark : ")
		fmt.Scan(&X)
		fmt.Print(currentPlayer.Name," \nEnter Y cordinate nuofmber where you wish to mark : ")
		fmt.Scan(&Y)

		icell := b.NCells[X][Y]
		if icell.GetMark()==cell.BattleShip{
			icell.SetMark(cell.Hit)
			b.Display()
			fmt.Println("Hurray! You hit the ship")
			b.DisplayHitMiss()
			shipCells--

		}else if icell.GetMark()==cell.NoMark{
			icell.SetMark(cell.Miss)
			fmt.Println("Oh no,you missed the ship")
			b.DisplayHitMiss()
		}

	}

	fmt.Println("GAME OVER!\n\n")
	b.Display()
	return
}
