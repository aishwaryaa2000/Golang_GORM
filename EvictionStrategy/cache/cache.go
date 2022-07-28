package cache

import "fmt"

type Algo interface {
	Evict(*Cache)
	Add(string)
}

type Cache struct {
	Storage     map[string]int
	algo        Algo
	size        int
	maxCapacity int
}

func NewCache(e Algo) *Cache {
	tempstorage := make(map[string]int)
	return &Cache{
		Storage:     tempstorage,
		algo:        e,
		size:        0,
		maxCapacity: 3,
	}
}

func (c *Cache) setEvictionAlgo(e Algo) {
	c.algo = e
}

func (c *Cache) Add(key string) {
	if c.size == c.maxCapacity {
		c.evict()
	}
	c.size++
	c.Storage[key] = 0
	c.algo.Add(key)
}

func (c *Cache) Get(key string) {
	_,okBool:= c.Storage[key]
	if !okBool{
		fmt.Println("Element doesn't exist")
	}else{
	c.Storage[key]++
	fmt.Println("Key:", key, " Value:", c.Storage[key])
	c.algo.Add(key)
	}
}

func (c *Cache) evict() {
	c.algo.Evict(c)
	c.size--
}
