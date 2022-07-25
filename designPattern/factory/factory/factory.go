package factory

import(
	 "fmt"
	 "factoryDp/product"

)

func GetFastFood(fastFoodType string) (product.Food, error) {
    if fastFoodType == "BelgianWaffle" {
        return product.NewWaffle(), nil
    }
    if fastFoodType == "CheeseSandwich" {
        return product.NewSandwich(), nil
    }
    return nil, fmt.Errorf("Wrong fast food type passed")
}
