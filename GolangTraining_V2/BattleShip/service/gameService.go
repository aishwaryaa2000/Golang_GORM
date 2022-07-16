package service

import (
	"battleShip/components/board"
	"battleShip/components/cell"
	"battleShip/components/player"
	"fmt"
)

func GameStart(b *board.Board, currentPlayer *player.Player) {
	var X, Y uint8
	var (
		shipCells, noOfTries uint = 15, 50
	)
	rowSize, colSize := b.GetRowColSize()
	fmt.Println("You only have ", noOfTries, " to attack all the ships")
	for shipCells > 0 && noOfTries > 0 {

		X, Y = inputFromPlayerToAttack(currentPlayer.GetName(), rowSize, colSize)

		icell := b.NCells[X][Y]

		if icell.Cell() == cell.BattleShip {
			hitTheShip(icell, b, currentPlayer, &shipCells, &noOfTries)
		} else if icell.Cell() == cell.NoMark {
			missedTheShip(icell, b, currentPlayer, &noOfTries)
		} else {
			fmt.Println("Already attacked this ship")
		}

		fmt.Println("Number of hits : ", currentPlayer.GetHit(), " \nNumber of miss : ", currentPlayer.GetMiss(), "\nTrials left : ", noOfTries)
	}

	resultAnalysis(shipCells, noOfTries)

}

func resultAnalysis(shipCells, noOfTries uint) {
	fmt.Println("GAME OVER")
	if shipCells == 0 {
		fmt.Println("You won the game!\nYou needed ", 50-noOfTries, " of tries to win the game")
	} else {
		fmt.Println("Oh no!You have exhausted all your tries")
	}
}

func hitTheShip(icell *cell.Cell, b *board.Board, currentPlayer *player.Player, shipCells *uint, noOfTries *uint) {
	fmt.Println("Hurray! You hit the ship")
	icell.SetMark(cell.Hit)
	b.DisplayHitMiss()
	currentPlayer.IncrementHit()
	*shipCells--
	*noOfTries--

}

func missedTheShip(icell *cell.Cell, b *board.Board, currentPlayer *player.Player, noOfTries *uint) {
	fmt.Println("Oh no,you missed the ship")
	icell.SetMark(cell.Miss)
	b.DisplayHitMiss()
	currentPlayer.IncrementMiss()
	*noOfTries--
}
