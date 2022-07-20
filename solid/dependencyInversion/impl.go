package main

import "fmt"

type Logger interface {
	Info()
	Debug()
}

type DefaultLogger struct {
	typeOfError int
}

type Zerolog struct {
	errorMsg string
}

type ActualLogger struct {
	Logger
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
	log := ActualLogger{DefaultLogger{}, "Golang"}
	log.Info()
	log.Debug()

	log = ActualLogger{Zerolog{}, "Java"}
	log.Info()
	log.Debug()

}