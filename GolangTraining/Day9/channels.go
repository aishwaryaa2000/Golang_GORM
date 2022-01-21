package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Inside main")
	ch := make(chan int) //read or write integer data on channel
	go writeToChannel(ch)

	// main() func doesn't wait for above  go routine to be completed so we need to introduce a delay
	// reading from a channel will cause delay
	// read will wait until some data is written
	// as soon as a data is written in the channel this reading part will get executed and main() will end
	value := <-ch //read from channel so data stored inside value
	fmt.Println("Data in our channel is : ",value)

}

func writeToChannel(ch chan int) {
	fmt.Println("Inside go routine")
	ch <- 10 //as soon as data is written in the channel,the main() resumes it's execution i.e delay is over
	time.Sleep(time.Second)  //causes delay so next statement is not executed at all
	fmt.Println("Data written to channel")
}


// channel provide medium of communication for go routines
// write to a channel or read from a channel