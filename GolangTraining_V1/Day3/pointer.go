package main
import "fmt"
import "reflect"

func main(){
	var a int =10
	fmt.Println("Value of a : ",a)
	fmt.Println("Memory address of a : ",&a)
	var p *int = &a //pointer
	fmt.Println("Value of p : ",p)
	fmt.Println("Value of a by derefencing : ",*p)
	fmt.Println("Type of p : ",reflect.TypeOf(p))
	fmt.Println("Memory address of p : ",&p)

	// pointer holding memory address of another pointer

	var pointerToPointer **int = &p
	// **int where first * indicates it's pointer 
	// and *int indicates that it pointer to a integer pointer

	fmt.Println("Value of p2p : ",pointerToPointer)
	fmt.Println("Value of p2p by derefencing : ",*pointerToPointer)
	fmt.Println("Double derefencing : ",**pointerToPointer)

	fmt.Println("Type of p2p : ",reflect.TypeOf(pointerToPointer))
	fmt.Println("Memory address of p2p : ",&pointerToPointer)


}