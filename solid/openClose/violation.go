package main

import "fmt"

type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (c circle) area() float64 {
	return 22 / 7 * (c.radius * c.radius)
}

func (s square) area() float64 {
	return s.length * s.length
}



func (c circle) vol() float64 {
	return 22 / 7 * (c.radius * c.radius * c.radius)
}

func (s square) vol() float64 {
	return s.length * s.length*s.length
}


func SumAreas(i ...interface{}) float64 {
	var sumArea float64 = 0
	for _, value := range i {
		switch value.(type) {
		case square:
			sqArea := value.(square).area()
			sumArea = sumArea + sqArea
		case circle:
			cirArea := value.(circle).area()
			sumArea = sumArea + cirArea

		}

	}

	return sumArea

}

func SumVol(i ...interface{}) float64 {
	var sumVol float64 = 0
	for _, value := range i {
		switch value.(type) {
		case square:
			sqVol := value.(square).vol()
			sumVol = sumVol + sqVol
		case circle:
			cirVol := value.(circle).vol()
			sumVol = sumVol + cirVol

		}

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


	
	ans = SumVol(s1, s2, c1, c2)
	fmt.Println("Ans for sum of volume : ",ans)
	//slice ka interface

}
