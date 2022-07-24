package featuredecorator

import (
	"decorator/watch"
)
type CallFunction struct {
    W watch.SmatWatch
}

func (c *CallFunction) GetPrice() int {
    watchPrice := c.W.GetPrice()
    return watchPrice + 2000
}
type SportsMode struct {
    W watch.SmatWatch
}
func (s *SportsMode) GetPrice() int {
    watchPrice := s.W.GetPrice()
    return watchPrice + 1000
}

