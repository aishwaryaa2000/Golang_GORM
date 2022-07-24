package main

import "fmt"

import (
	"decorator/watch"
	"decorator/featuredecorator"

)
func main() {

    watch := &watch.BoatSmartWatch{}

    watchWithCallFeature := &featuredecorator.CallFunction{
        W: watch,
    }

    watchWithCallFeatureSportsMode := &featuredecorator.SportsMode{
        W: watchWithCallFeature,
    }

    fmt.Printf("Price of smartwatch with call function and sports mode is %d\n", watchWithCallFeatureSportsMode.GetPrice())
}