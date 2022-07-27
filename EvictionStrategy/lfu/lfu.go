package lfu

import (
	"fmt"
	"math"
	"EvictionStrategy/cache"
)

type Lfu struct {
	Freq map[string]int
}

func NewLfu() *Lfu {
	return &Lfu{
		Freq: make(map[string]int),
	}
}

func (l *Lfu) Evict(c *cache.Cache) {
	fmt.Println("Evicting by lfu strtegy")
	var key string
	min := math.MaxInt
	for key1, val := range l.Freq {
		if min > val {
			min = val
			key = key1
		}
	}
	delete(l.Freq, key)
	delete(c.Storage, key)
}

func (l *Lfu) Add(key string) {
	l.Freq[key]++
}
