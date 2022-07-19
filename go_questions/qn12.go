package main

import (
	"fmt"
	"sync"
)

var wg = &sync.WaitGroup{}

func main() {
	wg.Add(1)

	nums := []int{1, 2, 3, 4, 5, 6}
	for _, num := range nums {
		// You can add code but can't remove to get the expected output
		// Expected output should be 1 2 3 4 5 6

		go func() {
			test(num)

		}()

	}


}

func test(i int) {
	fmt.Printf("%d ", i)
}
