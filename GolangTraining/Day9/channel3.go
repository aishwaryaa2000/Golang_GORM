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
	ch := make(chan int,2)
	ch2 := make(chan struct{}) //signal channel 
	go readFromChannel(ch)
	go writeToChannel(ch)
	wg.Wait()
	close(ch) //once channel is closed,we can't read or write on that channel
	go func () {
		ch2 <- struct{}{} //0 memory allocation signaling
	}()

	for i := range ch{   
		fmt.Println("Inside range",i)  
	}

	for i:=0;i<2;i++{
		select{
		case value,ok := <-ch:
			fmt.Println("Channel 1 has data ",value,ok)
		case value,ok := <-ch2:
			fmt.Println("Channel 2 has data ",value,ok)
		}

	}

}

func writeToChannel(ch chan<- int) { //send only channel means only writing on channel can be done in this func
	fmt.Println("Starting to write")
	time.Sleep(time.Second)  
	ch <- 10 
	ch <- 8 
	ch <- 6 //although 6 is written but not read ,deadlock doesn't occur as channel capacity is 2
	fmt.Println("Data written to channel")
	wg.Done()
}

func readFromChannel(ch <-chan int){
	fmt.Println("Started to read")
	value := <-ch
	value2 := <-ch
	fmt.Println("Data in our channel is : ",value,value2)
	wg.Done()

}


