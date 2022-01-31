package player
import(
	"tic_tac_toe_app/components/cell"
)
type Player struct {
	Mark cell.Mark //player's mark
	Name string
}

func New(playerMark cell.Mark,name string) *Player {
	var playerTest = &Player{
		Mark: playerMark,
		Name: name,
	}
	return playerTest //pointer to player
}