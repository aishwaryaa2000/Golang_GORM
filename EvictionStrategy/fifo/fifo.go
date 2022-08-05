package fifo

import (
	"fmt"
	"swabhav-training/GoTraining/EvictionStrategy/cache"
)

type Fifo struct {
}

func New() *Fifo {
	return &Fifo{}
}

func (l *Fifo) Evict(c *cache.Cache) {
	fmt.Println("Evicting by fifo strtegy")
	tempKey := c.AddOrder[0]
	delete(c.Storage, tempKey)
	c.AddOrder = c.AddOrder[1:]
}
