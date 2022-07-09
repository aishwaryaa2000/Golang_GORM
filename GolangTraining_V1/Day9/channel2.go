package main

import (
	"fmt"
	"sync"
	"time"
)
var wg = sync.WaitGroup{}

func main() {
	wg.Add(2)
	fmt.Println("Inside main")
	// this is a buffered channel with channel capacity 2
	ch := make(chan int,2) //channel capacity=2 means at max 2 data can be written but not read - no deadlock will occur
	go readFromChannel(ch)
	go writeToChannel(ch)
	wg.Wait()
	// main() func doesn't wait for above  go routine to be completed so we need to introduce a delay
	// reading from a channel will cause delay
	// read will wait until some data is written in channel
	// as soon as a data is written in the channel this reading part will get executed and main() will end
	// value := <-ch //read from channel so data stored inside value
	// fmt.Println("Data in our channel is : ",value)

}

func writeToChannel(ch chan int) {
	fmt.Println("Starting to write")
	time.Sleep(time.Second)  
	ch <- 10 //as soon as data is written in the channel,the main() resumes it's execution i.e delay is over
	ch <- 8 
	ch <- 6 //although 6 is written but not read ,deadlock doesn't occur as channel capacity is 2
	fmt.Println("Data written to channel")
	wg.Done()
}

func readFromChannel(ch chan int){
	fmt.Println("Started to read")
	value := <-ch
	value2 := <-ch
	fmt.Println("Data in our channel is : ",value,value2)
	wg.Done()

}
// channel provide medium of communication for go routines
// write to a channel or read from a channel

// writing is done first then reading