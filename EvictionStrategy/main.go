package main

import (
	"fmt"
	"swabhav-training/GoTraining/EvictionStrategy/cache"
	"swabhav-training/GoTraining/EvictionStrategy/lfu"
	"swabhav-training/GoTraining/EvictionStrategy/lru"
)

func main() {
	lfu := lfu.New()
	//fifo := fifo.New()
	lru := lru.New()
	cache := cache.New(lfu)

	cache.Add("a")
	cache.Add("b")

	cache.Add("c")

	cache.Get("b")
	cache.Get("b")
	cache.Get("a")
	cache.Get("a")
	cache.Get("c")
	cache.Get("a")

	cache.SetEvictionAlgo(lru)

	cache.Add("d")

	fmt.Println(cache.Storage)

	// lru := lru.NewLru()
	// cache1 := cache.NewCache(lru)

	// cache.Add("d")

	// fifo := fifo.NewFifo()
	// cache = cache.NewCache(fifo)
	// cache.Add(fifo)

	// cache.Add("e")

}
