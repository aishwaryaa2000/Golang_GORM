package cell

type Mark string
const (
	NoMark Mark = "   "
	Hit Mark  = " X "
	Miss Mark   = " M "
	BattleShip Mark   = " B "
)

type Cell struct {
	cellMark Mark

}

func New() *Cell {
	var cellTest = &Cell{
		cellMark: NoMark,
	}
	return cellTest //pointer to cell
}

func (c*Cell) Cell() (Mark){
	return c.cellMark
}

func (c *Cell) SetMark(markByUser Mark) {
	
		c.cellMark = markByUser
	
}


