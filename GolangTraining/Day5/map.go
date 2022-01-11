package main

import "fmt"

func main() {
	var myMap = make(map[string]int)
	//key is string data type and value is integer type
	myMap["a"] = 10
	myMap["b"] = 20
	myMap["c"] = 30

	fmt.Println(myMap)

	for key,value := range myMap{
		fmt.Println("Key is : ",key," Value is :",value)
	}

	delete(myMap,"c")
	fmt.Println(myMap)

	var num1,ok = myMap["a"]
    fmt.Println("\nValue of num1 : ",num1)
	fmt.Println("Ok value is : ",ok)

	num2,ok := myMap["C"]
	fmt.Println("\nValue of num2 : ",num2)
	fmt.Println("Ok value is : ",ok)
}

//ok value is true if key exists in map and false when key doesn't exist in map