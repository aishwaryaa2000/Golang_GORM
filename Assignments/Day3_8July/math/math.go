package math

func Add(no1,no2 float32) float64{
	return float64(no1 + no2)
}

func Subtract(no1,no2 float32) float64{
	return float64(no1 - no2)
}

func Mul(no1,no2 float32) float64{
	return float64(no1 * no2)
}

func Div(no1,no2 float32) float64{
	return float64(no1 / no2)
}

func MathOperations(num1,num2 float32,f func(float32,float32)float64)float64{
	return f(num1,num2)
}