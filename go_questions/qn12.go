package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = &sync.WaitGroup{}

func main() {

	nums := []int{1, 2, 3, 4, 5, 6}
	wg.Add(len(nums))
	for _, num := range nums {
		// You can add code but can't remove to get the expected output
		// Expected output should be 1 2 3 4 5 6

		go func() {
			test(num)
		}()
		time.Sleep(time.Second)
	}
	wg.Wait()
}

func test(i int) {
	fmt.Printf("%d ", i)
	wg.Done()
}