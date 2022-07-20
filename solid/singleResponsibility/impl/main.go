/* 
The single responsibility principle suggests that two separate aspects of a problem need to be handled by a different module.
In other word, changes in a module should be originated from only one reason.

circle has it's own module/package with area as it's method
and square has it's own module/package with area as it's method

if you have more than one responsibilities embedded into a single class, 
the internal logics become highly coupled, which makes the class less responsive towards changes
 
if printing or display of area is done in the method area then area() method is performing 2 operations-
1) calculate area
2) print area
So it is not a good practice

Instead have a Display() function that handles only the display part
Here,each function has it's own responsibility-
Area to calculate area 
Display function either displays in json format or simply in fmt.println or write onto file

*/

package main

import(
	"singleRes/shape/circle"
	"singleRes/shape/square"
	"singleRes/display"


)

func main(){
	c1 := circle.New(2)
	s1 := square.New(4)

	/* 
	In order to support polymorphism and abstraction,an interface shape can be created which implements Area() method.
	Then create a function that takes interface as parameter and executes the Area() method of the object that is called.
	So,w are simply passing the object but behind the scene,which Area() method is getting implemented is unknown to the user
	*/

	circleArea := c1.Area()
	squareArea := s1.Area()


	//DisplayAns function just takes up the responsibility of displaying the answer in json format or onto terminal or onto ans.txt
	display.DisplayAns(circleArea,"json")
	display.DisplayAns(squareArea,"terminal")
	display.DisplayAns(squareArea,"file")



	
}