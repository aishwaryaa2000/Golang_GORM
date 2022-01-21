package main
import(
	"tic_tac_toe_app/components/board"
)

func main(){
	boardn:=board.NewBoard()
	board.Display(boardn)

}