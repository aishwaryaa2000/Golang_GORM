package main

import (
	"fmt"
	"oops/company"
)
func main(){

	forcepoint := company.New() //pointer

	//Encapsulation
	fmt.Println(forcepoint.CeoName)// exported 
	fmt.Println("Number of employees : ",forcepoint.GetNoOfEmp())//exported
	// forcepoint.budget is unexported field


	//inheritance as address is embedded inside company
	forcepoint.PostalCode="400706"
	forcepoint.Street="tthanefp"
	forcepoint.City="navimum"
	forcepoint.State="maharashtra"


	fmt.Println("Address is : ",forcepoint.Address)

	/*
		Abstraction and ploymorphism

		PassInterface is a single function but is calling methods of different objects so this is polymorphism
		Poly means many and morphs means forms
		so the PassInterface has many forms depending upon the object which is passed
		
		The internal working is hidden,i.e which GetName() is called behind the PassInterface is hidden from the user
	*/
	teamRbi := company.NewRbi()
	company.PassInterface(teamRbi)

	teamNetsec := company.NewNetSec()
	company.PassInterface(teamNetsec)

	var a company.Address 
	a.Street="Thane"
	a.PostalCode="4007066"
	a.City="navi mumbai"
	a.State="maha"

	// company.ParameterWithLargerStruct(a)
	//you cannot pass smaller struct when parameter is larger struct

	// company.ParameterWithSmallerStruct(*forcepoint) //derefencing
	//you cannot pass larger struct into parameter with smaller struct
}

// interface creation is abstraction
// interface use karna is polymorphism because same function is called but internally different methods are executed