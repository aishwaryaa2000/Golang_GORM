package cell
import "errors"
type Mark string
const (
	NoMark Mark  = "-"
	XMark Mark   = "X"
	OMark Mark   = "O"
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

func (c*Cell) GetMark() (Mark){
	return c.cellMark
}

func (c *Cell) SetMark(markByUser Mark) (error) {
	if c.cellMark == NoMark {
		c.cellMark = markByUser
		return  nil
	}
	return errors.New("This cell is already marked")
}
