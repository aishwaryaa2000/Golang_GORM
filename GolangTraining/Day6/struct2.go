package main
import "Day6/structure"
import "fmt"
func main(){
	var person1 = structure.NewPerson("Sahil","Chavan",20)
	fmt.Println(person1)
	fmt.Println(person1.FirstName)
	// fmt.Println(person1.lastName) error bcoz it is private member or unexported

}