package main
import "fmt"

func main(){

	fmt.Println("FOR CIRCLE")
	passInterface(circle{ //Notice no &circle
		radius: 10.,
	})

	fmt.Println("\nFOR RECTANGLE")

	passInterface(&rectangle{ //Notice &rectangle
		length: 10,
		breadth: 20,
	})
	
}

type shape interface{
	area() float64 //methods
	perimeter() float64 //method
}

type circle struct{
	radius float64
}

type rectangle struct{
	length float64
	breadth float64
}

//struct circle implements shape interface bcoz circle implements these two methods
func(c circle)area() float64{  //notice no c *cirle
	return 22/7*(c.radius*c.radius)
}

func(c circle)perimeter() float64{
	return 2*22/7*c.radius
}

//struct rectangle implements shape interface bcoz rectangle implements these two methods
func(r *rectangle)area() float64{ //notice r*rectangle
	return r.length*r.breadth
}

func(r rectangle)perimeter() float64{ //notice no r*rectangle
	return 2*(r.breadth+r.length)
}


func passInterface(sh shape){ 
	//a func that needs shape interface accepts circle or rectangle as argument bcoz circle&rectangle implements shape
	fmt.Println("Area is: ",sh.area())
	fmt.Println("Perimeter is : ",sh.perimeter())
}

//every method of interface shape MUST BE IMPLEMENTED by the circle or rectangle structure ELSE ERROR

/* If in argument we pass value 
passInterface(circle{
		radius: 10.,
	})
and in the parameter we have pointer receiver 
func(c *circle)area() float64{
	return 22/7*(c.radius*c.radius)
}
 THEN ERROR


 BUT IF
passInterface(circle{
		radius: 10.,
	})
and in the parameter we have pointer receiver 
func(c circle)area() float64{
	return 22/7*(c.radius*c.radius)
}
NO ERROR

passInterface(&circle{
		radius: 10.,
	})
and in the parameter we have pointer receiver 
func(c circle)area() float64{
	return 22/7*(c.radius*c.radius)
}
NO ERROR
 But value receiver work on both pointer as well as value
*/