package main

import "fmt"

func main() {
	var person1 = &person{ //person1 is a pointer now
		firstName : "Aishwarya",
		age : 30,
	}
	fmt.Println(person1) //automatically dereferencing of pointer takes place
	fmt.Println(person1.age) //automatically dereferencing of pointer takes place
	fmt.Println("Address of structure",&person1)
	fmt.Println("Value at structure pointer",*person1)

	var person2 = newPerson("Aayush","Anand",12)
	fmt.Println(*person2)
}

func newPerson(firstName,lastName string,age int) *person{
	var personTest = &person{
		firstName: firstName, //fieldName : VariableName
		lastName: lastName,
		age: age,
	}
	return personTest //this is a pointer
}

type person struct {
	firstName string
	lastName string
	age       int
}