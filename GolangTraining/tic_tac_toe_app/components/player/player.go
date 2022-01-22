package player
import(
	"tic_tac_toe_app/components/cell"
)
type player struct {
	mark cell.Mark //player's mark
	name string
}

func New(playerMark cell.Mark) *player {
	var playerTest = &player{
		mark: playerMark,
		// name: "", means name is empty string
	}
	return playerTest //pointer to player
}