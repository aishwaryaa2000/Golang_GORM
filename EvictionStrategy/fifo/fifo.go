package fifo

import (
	"fmt"
	"EvictionStrategy/cache"
)

type Fifo struct {
	Data []string
}

func NewFifo() *Fifo {
	return &Fifo{
		Data: make([]string, 0),
	}
}

func (l *Fifo) Evict(c *cache.Cache) {
	fmt.Println("Evicting by fifo strtegy")
	tempKey := l.Data[0]
	delete(c.Storage, tempKey)
	l.Data = l.Data[1:]

}

func (l *Fifo) Add(key string) {
	l.Data = append(l.Data, key)
}
