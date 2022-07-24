/*
END GOAL is to find the area.
But there are different algorithms or strategies to find area of a triangle,each of this algorithm is a different structure implementing
Area{} interface.

Strategy pattern suggests creating a family of the algorithm with each algorithm having its own class.Example-HeronsArea,PrimaryArea

Each of these classes follows the same interface Area{}, and this makes the algorithm interchangeable within the family.
Let’s say the common interface name is Area.

*/

package main

import (
	"fmt"
	"math"
)


type Triangle struct{
	//context
	aSide,base,cSide,height float64
	area AreaAlgorithm

}

func (t *Triangle) setAreaAlgorithm(algorithmByUser AreaAlgorithm) {
	//Setting which strategy to use- setStrategy = setAreaAlgorithm
    t.area = algorithmByUser
}

func (t *Triangle)executeStrategy(){
	/*The context delegates some work to the strategy object
    instead of implementing multiple versions of the
    algorithm on its own.*/
	t.area.findArea(t)
}


type AreaAlgorithm interface {
    findArea(t *Triangle) 
}

type HeronsArea struct{
	//Strategy 1
}
func (h *HeronsArea) findArea(t *Triangle) {
	semiPerimeter := (t.aSide+t.base+t.cSide)/2
	area:= math.Sqrt(semiPerimeter*(semiPerimeter-t.aSide)*(semiPerimeter-t.base)*(semiPerimeter-t.cSide))
	fmt.Println("Area of triangle using Heron's formula  is : ",area)
}

type PrimaryArea struct{
	//Strategy 2
}
func (p *PrimaryArea) findArea(t *Triangle) {
	area := (t.base*t.height)/2
	fmt.Println("Area of triangle using primary formula of 1/2*base*height is : ",area)

}


func main(){
	triangle:= Triangle{
					aSide:3,
					base:4,
					cSide : 5,
					height: 3,
				}
	triangle.setAreaAlgorithm(&HeronsArea{})
	triangle.executeStrategy()

	triangle.setAreaAlgorithm(&PrimaryArea{})
	triangle.executeStrategy()
}