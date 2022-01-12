package structure

func NewPerson(firstName,lastName string,age int) *person{
	var personTest = &person{
		FirstName: firstName, //fieldName : VariableName
		lastName: lastName,
		Age: age,
	}
	return personTest //this is a pointer
}

type person struct {
	FirstName string
	lastName string
	Age       int
}


// import "fmt"

// func main() {
// 	var person1 = &person{ //person1 is a pointer now
// 		firstName : "Aishwarya",
// 		age : 30,
// 	}
// 	fmt.Println(person1) //automatically dereferencing of pointer takes place
// 	fmt.Println(person1.age) //automatically dereferencing of pointer takes place
// 	fmt.Println("Address of structure",&person1)
// 	fmt.Println("Value at structure pointer",*person1)

// 	var person2 = newPerson("Aayush","Anand",12)
// 	fmt.Println(*person2)

// 	var Employee1 = employee{
// 		employeeID: 100,
// 		person: person{
// 			firstName: "Ramani",
// 			lastName: "Anand",
// 			age: 45,
// 		},
// 	}

// 	fmt.Println(Employee1)
// 	fmt.Println("Employee ID is : ",Employee1.employeeID)
// 	fmt.Println("Employee Name is : ",Employee1.person.firstName)




// }



// type employee struct{
// 	employeeID int
// 	person //this means employee IS also a person
// }