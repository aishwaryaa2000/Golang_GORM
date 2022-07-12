package player

type Player struct {
	Name string
}

func New(name string) *Player {
	var playerTest = &Player{
		Name: name,
	}
	return playerTest //pointer to player
}