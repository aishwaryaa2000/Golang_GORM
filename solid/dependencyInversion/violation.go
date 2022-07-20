package main

import "fmt"


type DefaultLogger struct {
	typeOfError int
}

type Zerolog struct {
	errorMsg string
}

type ActualLogger struct {
	Zerolog //We are predefined that zerolog is a member of actual logger but what if we default logger at times then we need to change the member here again and again
	packageName string
}

func (d DefaultLogger) Info() {
	fmt.Println("Hey!I am default logger")
}

func (d DefaultLogger) Debug() {
	fmt.Println("I debug golang error using default logger")
}

func (z Zerolog) Info() {
	fmt.Println("Hey!I am zerlogger dedicated library for structured JSON logging")
}

func (z Zerolog) Debug() {
	fmt.Println("I debug golang or json error using zerlog logger")
}

func main() {
	log := ActualLogger{Zerolog{errorMsg:"yo",}, "Golang"}
	log.Info()
	log.Debug()


}