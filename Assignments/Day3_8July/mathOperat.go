package main
import(
	"fmt"
	"assignments/Day3_8July/math"
)
func main(){
	var num1 float32 =8.0
	var num2 float32 =2.0
	sum := math.MathOperations(num1, num2, math.Add)
	product := math.MathOperations(num1, num2, math.Mul)
	diff := math.MathOperations(num1, num2, math.Subtract)
	quotient := math.MathOperations(num1, num2, math.Div)

	fmt.Println("Sum of",num1," and ",num2 ," is :", sum)
	fmt.Println("Difference between",num1," and ",num2 ," is :", diff)
	fmt.Println("Product of of",num1," and ",num2 ," is :", product)
	fmt.Println("Quotient of of",num1," / ",num2 ," is :", quotient)

}