package service

import (
	"battleShip/components/board"
	"battleShip/components/cell"
	"battleShip/components/player"
	"fmt"
)

func checkIfShipHitOrMiss(icell *cell.Cell,b *board.Board, currentPlayer *player.Player, noOfTries ,shipCells *uint){

	markAtCell := icell.Cell()

	switch markAtCell{
	case cell.BattleShip:
		hitTheShip(icell, b, currentPlayer, shipCells, noOfTries)
	case cell.NoMark:
		missedTheShip(icell, b, currentPlayer, noOfTries)
	default:
		fmt.Println("Already attacked this ship")
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

func placeShip(b [][]*cell.Cell, min, max, cordinate int, direction rune) {
	//to place ships on a board
	var icell *cell.Cell
	for i := min; i < max; i++ {
		if direction == 'H' {
			icell = b[cordinate][i]
		} else {
			icell = b[i][cordinate]
		}
		icell.SetMark(cell.BattleShip)
	}

}