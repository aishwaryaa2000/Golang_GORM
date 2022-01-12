package main

import "fmt"

func main() {
	var person1 = person{
		firstName: "Aishwarya",
		lastName:  "Anand",
		age:       21,
		address: []string{
			"Mumbai",
			"Pune",
		},
	}
	var person2 = person{
		firstName: "Aayush",
		lastName:  "Anand",
		age:       13,
		address: []string{
			"Mumbai",
		},
	}
	var person3 = person{
		firstName: "Ramani",
		lastName:  "Anand",
		age:       51,
		address: []string{
			"Mumbai",
			"Pune",
		},
	}
	fmt.Println("Before calling modifyPerson : ",person1)
	modifyPerson(person1)
	fmt.Println("After calling modifyPerson : ",person1)
	modifyPersonByPointer(&person1)
	fmt.Println("After calling modifyPersonByPointer : ",person1)

	var allPerson[] person
	allPerson=[]person{person1,person2,person3}
	for _,singlePerson := range allPerson{
		fmt.Println("\nFirst Name is : ",singlePerson.firstName)
		fmt.Println("Last Name is : ",singlePerson.lastName)
		fmt.Println("Age is : ",singlePerson.age)
		
		for index,add := range singlePerson.address{
			fmt.Println("Address ",index+1,"is : ",add)

		}
	}

}
func modifyPerson(p person) {
	//passed by value.So changes inside won't be reflected in main()
	p.firstName = "Aishu"
	p.age = 50
	fmt.Println("Inside modifyPerson : ",p)
}
func modifyPersonByPointer(p *person) {
	//passed by pointer.So changes inside will be reflected in main()
	p.firstName = "Aishu"
	p.age = 50
	fmt.Println("Inside modifyPersonByPointer : ",p)
}

type person struct {
	firstName string
	lastName  string
	age       int
	address   []string
}