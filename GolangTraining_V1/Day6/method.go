package main

import "fmt"

func main() {
	var person1 = person{
		firstName: "Aishwarya",
		lastName:  "Anand",
		age:       21,
		address: []string{
			"Mumbai",
			"Navi Mumbai",
		},
	}
	fmt.Println(person1)
	fmt.Println(person1.getFullName())
	var person2 = person{
		firstName: "Ramesh",
		lastName:  "Yadav",
		age:       21,
	}
	fmt.Println(person2.getFullName())

}

func (p *person) getFullName() string{
	return p.firstName + " " + p.lastName
}
type person struct {
	firstName string
	lastName  string
	age       int
	address   []string
}