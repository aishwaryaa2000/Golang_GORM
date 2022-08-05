package lfu

import (
	"fmt"
	"math"
	"swabhav-training/GoTraining/EvictionStrategy/cache"
)

type Lfu struct {
}

func New() *Lfu {
	return &Lfu{}
}

func (l *Lfu) Evict(c *cache.Cache) {
	fmt.Println("Evicting by lfu strtegy")
	var key string
	min := math.MaxInt
	for key1, val := range c.Storage {
		if min > val {
			min = val
			key = key1
		}
	}
	delete(c.Storage, key)
	//Deleting from AddOrder
	c.DeleteFromAddOrder(key)
	//Deleting from ReadOrder
	c.DeleteFromReadOrder(key)
}
