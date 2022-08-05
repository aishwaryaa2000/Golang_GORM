package lru

import (
	"fmt"
	"swabhav-training/GoTraining/EvictionStrategy/cache"
)

type Lru struct {
}

func New() *Lru {
	return &Lru{}
}

func (l *Lru) Evict(c *cache.Cache) {
	fmt.Println("Evicting by lru strtegy")
	tempkey := c.ReadOrder[0]
	delete(c.Storage, tempkey)
	c.ReadOrder = c.ReadOrder[1:]
	c.DeleteFromAddOrder(tempkey)

}
