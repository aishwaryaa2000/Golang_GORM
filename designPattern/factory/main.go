package main

import(
	 "fmt"
	 "factoryDp/factory"
	 "factoryDp/product"

)

func main() {
    belgianWaffle, _ := factory.GetFastFood("BelgianWaffle")
    cheeseSandwich, _ := factory.GetFastFood("CheeseSandwich")

    printDetails(belgianWaffle)
    printDetails(cheeseSandwich)
}

func printDetails(f product.Food) {
    fmt.Printf("Fast food name is : %s\n", f.GetName())
    fmt.Printf("Calorie : %d \n", f.GetCalorie())
}