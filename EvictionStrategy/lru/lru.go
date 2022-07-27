package lru

import (
	"fmt"
	"EvictionStrategy/cache"
)

type Lru struct {
	Data []string
}

func NewLru() *Lru {
	return &Lru{
		Data: make([]string, 0),
	}
}

func (l *Lru) Evict(c *cache.Cache) {
	fmt.Println("Evicting by lru strtegy")
	tempkey := l.Data[0]
	delete(c.Storage, tempkey)
	l.Data = l.Data[1:]
}

func (l *Lru) Add(key string) {
	pos := -1
	for i, val := range l.Data {
		if val == key {
			pos = i
			break
		}
	}
	if pos != -1 {
		for i := pos; i < len(l.Data)-1; i++ {
			l.Data[i] = l.Data[i+1]
		}
		l.Data[len(l.Data)-1] = key
	} else {
		l.Data = append(l.Data, key)
	}
	fmt.Println(l.Data)
}
