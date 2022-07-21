package circle

type Circle struct {
	radius float64
}

func New(r float64) Circle{
	var CircleTest = Circle{
		radius: r,
	}
	return CircleTest
}
func (c Circle) Area() float64 {
	return 22 / 7 * (c.radius * c.radius)
}