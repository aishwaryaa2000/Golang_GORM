package company

import (
	"fmt"
)

type Company struct {
	budget              int
	CeoName             string
	noOfEmployeeWorking int
	Address
}
type Address struct {
	Street, City, State, PostalCode string
}

type RBI struct {
	nameOfManager string
}

type NetSec struct {
	nameOfManager string
}

func NewRbi() *RBI {
	var RBItest = &RBI{
		nameOfManager: "Ashwin Patel",
	}
	return RBItest //pointer to cell
}
func NewNetSec() *NetSec {
	var NetSecTest = &NetSec{
		nameOfManager: "Ashwin Patel",
	}
	return NetSecTest //pointer to cell
}

func (r RBI) GetName() string { //notice no c *cirle
	return r.nameOfManager
}

func (n NetSec) GetName() string {
	return n.nameOfManager
}

type team interface {
	GetName() string //methods
}

func PassInterface(t team) {
	fmt.Println("Manager is: ", t.GetName())
}

func New() *Company {
	var companyTest = &Company{
		budget:              10000,
		CeoName:             "Someone",
		noOfEmployeeWorking: 1999,
		// Street:              "streetthane",
		// City:                "Thane",
		// State:               "Maharashtra",
		// PostalCode:          "400706",
	}

	return companyTest //pointer to cell
}

func (c *Company) GetNoOfEmp() int {
	return c.noOfEmployeeWorking
}

func ParameterWithLargerStruct(c Company) {
	fmt.Println("You can or cannot pass smaller embedded struct with parameter in larger struct")
}

func ParameterWithSmallerStruct(a Address) {
	fmt.Println("You can or cannot pass larger embedded struct with parameter in smmaller struct")
}

/*
encapsulation me unexported and exported members
Capital lettered variables are accessible-exported
Lower letter variables are inaccessible-unexported
*/
