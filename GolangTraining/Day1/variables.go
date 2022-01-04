package main

import(
   "fmt"
   "runtime"
)


var name string //can be declared but can't be initialized here (error)

func main(){

   var rollNo int
   rollNo = 20
   rollNo = 30
   fmt.Println(rollNo)


   //if datatype is not specified then by default the type is considered acc to the value assigned
   var companyName="Forcepoint"
   fmt.Println(companyName)
  

   var(
 	firstName, lastName, flag = "Aishwarya","Anand",false
   )
   fmt.Println(firstName,flag,lastName)

   name="Aishwarya"
  
   test:= 20
   fmt.Println(test)
   fmt.Printf("%T",test) //to display datatype


   fmt.Println(runtime.GOOS)
   fmt.Println(runtime.GOARCH)

}


//If variable is declared outside function,it has to be initialized inside the function and not outside the function(error).
//If variable is declared outside function,it is not mandatory to use it.
//But if variable is declared inside function,then it MUST be used else error.
//If we try to update test:=20 by test:="abc" ,then it will raise an error as datatype of 'test' can't be changed from int to string.
//runtime package to know about our OS.



