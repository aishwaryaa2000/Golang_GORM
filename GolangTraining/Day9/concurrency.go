package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	
	runtime.GOMAXPROCS(1) //it will only use one thread although there are 2 go routines

	wg.Add(2) //2 go routines so wg.Wait() is expecting 2 wg.Done()
	go delayedIteration1()
	go delayedIteration2()
	wg.Wait() //wait until done is encountered
	// as soon as 2 wg.Done() is executed then program exits


	// main() func doesn't wait for above two go routines to be completed so we need to introduce this delay
	// fmt.Scanln() 

}

func delayedIteration1() {
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second)//wait for 1 second
		fmt.Println("In 1 Second : ",i)
	}

	wg.Done()

}


func delayedIteration2() {
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second)//wait for 1 second
		fmt.Println("In 2 Second : ",i)
	}

	wg.Done()

}

//two go routines can be executed simultaneously
// func main() is also a go routine
