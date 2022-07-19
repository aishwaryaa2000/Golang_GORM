package main

import (
	"fmt"
)

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
	Address		Address

}

type Employee struct {
	Name   string
	Salary float64
	Address		Address
}
type Address struct{
	Street,City,State,PostalCode string

}

// YOUR CODE HERE:
// Define a struct type named Address that has Street, City, State,
// and PostalCode fields, each with a type of "string".
// Then embed the Address type within the Subscriber and Employee
// types using anonymous fields, so that the code in "main" will
// compile, run, and produce the output shown.

func main() {
	var subscriber Subscriber
	subscriber.Name = "Aman Singh"
	AddressSub := Address{
		Street: "123 Oak St",
		City: "Omaha",
		State: "NE",
		PostalCode: "68111",
	} 
	subscriber.Address=AddressSub
	
	fmt.Println("Name:", subscriber.Name)              // => Name: Aman Singh
	fmt.Println("Street:", subscriber.Address.Street)          // => Street: 123 Oak St
	fmt.Println("City:", subscriber.Address.City)              // => City: Omaha
	fmt.Println("State:", subscriber.Address.State)            // => State: NE
	fmt.Println("Postal Code:", subscriber.Address.PostalCode) // => Postal Code: 68111

	var employee Employee
	AddressEmp := Address{
		Street: "456 Elm St",
		City: "Portland",
		State: "OR",
		PostalCode: "97222",
	} 
	employee.Name = "Joy Carr"
	employee.Address=AddressEmp
	fmt.Println("Name:", employee.Name)              // => Name: Joy Carr
	fmt.Println("Street:", employee.Address.Street)          // => Street: 456 Elm St
	fmt.Println("City:", employee.Address.City)              // => City: Portland
	fmt.Println("State:", employee.Address.State)            // => State: OR
	fmt.Println("Postal Code:", employee.Address.PostalCode) // => Postal Code: 97222
}
