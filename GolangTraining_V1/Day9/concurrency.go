package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
//wg is used so that main func waits for other two go routines get executed

func main() {
	
	runtime.GOMAXPROCS(1) //it will only use one thread although there are 2 go routines

	wg.Add(2) //2 go routines so wg.Wait() is expecting 2 wg.Done(). 2 means it is expecting 2 Done() signals
	go delayedIteration1()
	go delayedIteration2()
	wg.Wait() //main thread is waiting until 2 done signals is encountered
	// as soon as 2 wg.Done() is executed then program exits


	// main() func doesn't wait for above two go routines to be completed so we need to introduce this delay
	// fmt.Scanln() 

}

func delayedIteration1() {
	//1 iteration per second -> 5 iterations
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second)//wait for 1 second and then print value of i
		fmt.Println("In func1 Second : ",i)
	}

	wg.Done()

}


func delayedIteration2() {
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second)//wait for 1 second
		fmt.Println("In func2 Second : ",i)
	}

	wg.Done()

}

//two go routines can be executed simultaneously
//func main() is also a go routine
//so 3 go routines

//main go routine is the boss,if main gets completed then main will not wait for other go routines.
