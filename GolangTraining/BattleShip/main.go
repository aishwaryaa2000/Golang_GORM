package main

import(
	"battleShip/components/board"
)

func main(){
boardGame := board.New(4,6)
boardGame.Display()

}
