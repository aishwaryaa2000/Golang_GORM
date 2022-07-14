package player

type Player struct {
	Name string
	Hit uint
	Miss uint
}

func New(name string) *Player {
	var playerTest = &Player{
		Name: name,
		Hit:0,
		Miss:0,
	}
	return playerTest //pointer to player
}