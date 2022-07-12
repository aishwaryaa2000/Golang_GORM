package main

import(
	"battleShip/components/board"
)

func main(){
boardGame := board.New(11,7)
boardGame.Display()
boardGame.BoardInit()
boardGame.Display()



}
