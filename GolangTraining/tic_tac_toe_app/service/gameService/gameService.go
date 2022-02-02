package gameService

import (
	"fmt"
	"os"
	"tic_tac_toe_app/components/board"
	"tic_tac_toe_app/components/cell"
	"tic_tac_toe_app/components/player"
	"tic_tac_toe_app/components/result"
)

func Player(name1,name2 string) (*player.Player,*player.Player){

	Player1 := player.New(cell.XMark,name1)
	Player2 := player.New(cell.OMark,name2)
	return Player1,Player2
}

func Board(size uint8) *board.Board{
	board := board.New(size)
	board.Display()
	result.MakeWinArray(int(size))
	return board
}

func GameStart(b *board.Board,player1,player2 *player.Player){
	var index int
	noOfTurns := 0
	currentPlayer := player1
	for{
		noOfTurns++
		errorOccured:
		fmt.Print(currentPlayer.Name," enter cell number where you wish to mark : ")
		fmt.Scan(&index)
		if index >= int(b.Size*b.Size) || index<0{
			fmt.Println("Error occurred : Cell number entered is wrong.It should be between 0 to",b.Size*b.Size-1)
			goto errorOccured
		}
		icell := b.NCells[index]
		errorSetMark := icell.SetMark(currentPlayer.Mark)
		if(errorSetMark!=nil){
			fmt.Println("Error occured : ",errorSetMark)
			goto errorOccured
		}
		b.Display()
		checkForStatus(b.Size,currentPlayer,index,noOfTurns)
		currentPlayer = playerSwitch(currentPlayer,player1,player2)
	}

}



func playerSwitch(currentPlayer,player1,player2 *player.Player) *player.Player{
	if(currentPlayer==player1){
		currentPlayer=player2
	}else{
		currentPlayer=player1
	}
	return currentPlayer
}

func checkForStatus(boardSize uint8,currentPlayer *player.Player,index ,noOfTurns int){
		gameStatus := result.GetStatus(int(boardSize),currentPlayer.Mark,index,noOfTurns)
		if(gameStatus==result.Win){
			fmt.Println("Yaay!",currentPlayer.Name," won")
			os.Exit(0)
		}else if(gameStatus==result.Tie){
			fmt.Println("Game tie")
			os.Exit(0)
		}
}