package main

import (
	"gorm/app"
	"gorm/router"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	router := router.CreateRoute()
	app := app.New(router)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	//os.Interrupt like CRTL + C
	go app.Start()
	<-done
	app.Stop()
	//https://medium.com/over-engineering/graceful-shutdown-with-go-http-servers-and-kubernetes-rolling-updates-6697e7db17cf

}
