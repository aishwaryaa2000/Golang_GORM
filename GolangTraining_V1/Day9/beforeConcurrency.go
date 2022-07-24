package main

import (
	"fmt"
	"time"
)


func main() {

	delayedIteration1()
	delayedIteration2()
	

}

func delayedIteration1() {
	//1 iteration per second -> 5 iterations
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second)//wait for 1 second and then print value of i
		fmt.Println("In func1 Second : ",i)
	}


}


func delayedIteration2() {
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second)//wait for 1 second
		fmt.Println("In func2 Second : ",i)
	}


}

/* 
This program will take 10 seconds.
5 seconds for delayedIteration1()
and 5 seconds for delayedIteration2()

first delayedIteration1() is executed
then delayedIteration2() is executed



to have concurrency then introduce go routines
*/
