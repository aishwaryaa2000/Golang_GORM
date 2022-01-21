package cell

import "errors"

type Cell struct {
	mark string
}

func NewCell() *Cell {
	var cellTest = &Cell{
		mark: "-",
	}
	return cellTest

}

func (c *Cell) markCell(mark string) (*Cell, error) {
	if c.mark == "-" {
		c.mark = mark
		return c, nil
	}
	return c, errors.New("This cell is already marked")
}
