package main

import (
	"fmt"
	"gorm/app"
	"gorm/connect"
	"gorm/model"
	"gorm/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)


func main() {

	router := router.CreateRoute()
	app.New(router)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go app.Start()
	fmt.Println("Server Started")

	<-done
	log.Print("Server Stopped")

}

