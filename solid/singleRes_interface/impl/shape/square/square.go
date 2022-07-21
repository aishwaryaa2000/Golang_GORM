package square
type Square struct {
	length float64
}

func New(len float64) Square{
	var SquareTest = Square{
		length: len,
	}
	return SquareTest
}
func (s Square) Area() float64 {
	return s.length * s.length
}
