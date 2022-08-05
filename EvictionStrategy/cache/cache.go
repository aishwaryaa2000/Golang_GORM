package cache

import "fmt"

type Evictor interface {
	Evict(*Cache)
}

type Cache struct {
	Storage     map[string]int
	AddOrder    []string
	ReadOrder   []string
	algo        Evictor
	size        int
	maxCapacity int
}

func New(e Evictor) *Cache {
	tempstorage := make(map[string]int)
	return &Cache{
		Storage:     tempstorage,
		algo:        e,
		size:        0,
		maxCapacity: 3,
	}
}

func (c *Cache) SetEvictionAlgo(e Evictor) {
	c.algo = e
}

func (c *Cache) Add(key string) {
	if c.size == c.maxCapacity {
		c.evict()
	}
	c.size++
	c.Storage[key] = 0
	c.AddOrder = append(c.AddOrder, key)
	c.ReadOrder = append(c.ReadOrder, key)
}

func (c *Cache) Get(key string) {
	c.Storage[key]++
	fmt.Println("Key:", key, " Value:", c.Storage[key])
	pos := -1
	for i, val := range c.ReadOrder {
		if val == key {
			pos = i
			break
		}
	}
	if pos != -1 {
		for i := pos; i < len(c.ReadOrder)-1; i++ {
			c.ReadOrder[i] = c.ReadOrder[i+1]
		}
		c.ReadOrder[len(c.ReadOrder)-1] = key
	} else {
		c.ReadOrder = append(c.ReadOrder, key)
	}
}

func (c *Cache) evict() {
	c.algo.Evict(c)
	c.size--
}

func (c *Cache) DeleteFromReadOrder(key string) {
	pos := -1
	for i, val := range c.ReadOrder {
		if val == key {
			pos = i
		}
	}
	c.ReadOrder[pos] = c.ReadOrder[len(c.ReadOrder)-1]
	c.ReadOrder = c.ReadOrder[:len(c.ReadOrder)-1]
}

func (c *Cache) DeleteFromAddOrder(key string) {
	pos := -1
	for i, val := range c.AddOrder {
		if val == key {
			pos = i
		}
	}
	c.AddOrder[pos] = c.AddOrder[len(c.AddOrder)-1]
	c.AddOrder = c.AddOrder[:len(c.AddOrder)-1]
}
