package main

import (
	"fmt"
	"EvictionStrategy/cache"
	"EvictionStrategy/lru"
)

func main() {
	//lfu := lfu.NewLfu()
	//fifo := fifo.NewFifo()
	lru := lru.NewLru()
	cache := cache.NewCache(lru)

	cache.Add("a")
	cache.Add("b")

	cache.Add("c")

	cache.Get("b")
	cache.Get("b")
	cache.Get("a")
	cache.Get("a")
	cache.Get("c")
	cache.Get("a")

	cache.Add("d")

	fmt.Println(cache.Storage)
}
