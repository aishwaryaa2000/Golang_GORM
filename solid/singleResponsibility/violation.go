package main

import (
	"fmt"
)

type circle struct{
	radius float64
}

type square struct{
	length float64
}

func(c circle)area() float64{  
	return 22/7*(c.radius*c.radius)
}

func(s square)area() float64{
	return s.length*s.length
}



func main(){

	
	var c circle
	c.radius=2.2

	/* 
	c := circle{ 
     radius: 2.2,
	}
	*/

	areaValue := c.area()

	fmt.Println("Area of circle : ",areaValue)
	var s square
	s.length=4
	areaValue = s.area()
	fmt.Println("Area of square : ",areaValue)

	
}


/* 
If output is required in different format then we have to change the code a lot so violated
display (format)
if format is json or console then print the output accordingly
*/


