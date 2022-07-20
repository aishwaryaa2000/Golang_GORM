/* 
Open close
functionality should be open for extension and closed for modification
*/


package main

import "fmt"

type circle struct {
	radius float64
}

type square struct {
	length float64
}

type Twoshape interface{
	area() float64
}

type ThreeDShape interface{
	Twoshape
	vol() float64
}

func (c circle) area() float64 {
	return 22 / 7 * (c.radius * c.radius)
}

func (c circle) vol() float64 {
	return 22 / 7 * (c.radius * c.radius * c.radius)
}

func (s square) area() float64 {
	return s.length * s.length
}





func SumAreas(i ...Twoshape) float64 {
	var sumArea float64 = 0
	for _, value := range i {
		sumArea=sumArea +  value.area()
	}

	return sumArea

}

func SumVol(i ...ThreeDShape) float64 {
	var sumVol float64 = 0
	for _, value := range i {
		sumVol=sumVol + value.vol()
		}


	return sumVol

}

func main() {

	var c circle
	c.radius = 2.2

	c1 := circle{
		radius: 2.2,
	}

	c2 := circle{
		radius: 1.2,
	}

	s1 := square{
		length: 2,
	}

	s2 := square{
		length: 4,
	}

	ans := SumAreas(s1, s2, c1, c2)
	fmt.Println("Ans for sum of area : ",ans)


	
	ans = SumVol( c1, c2)

	//if i add s1,s2,c1,c2,cube1 then we have to change the switch case 
	fmt.Println("Ans for sum of volume : ",ans)
	//slice ka interface

}
