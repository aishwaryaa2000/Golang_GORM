package main
import(
	"fmt"
)

// Solve this without running the code
func main() {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3, 4)
	a := append(s, 5)
	fmt.Println(a) // O/P :- [1,2,3,4,5]
	b := append(s, 6) //s ka value is 1 2 3 4  but b ka value is 1 2 3 4 6
	fmt.Println(b) // O/P :- [1,2,3,4,6] 
	fmt.Println(a) // O/P :- [1,2,3,4,6]
	fmt.Println(s) // O/P :- [1,2,3,4]

	// What is the expected output and why?
	//whenever we append and assign then reference is used i.e address value is affected 
	//so while printing a then last index is 6 instead of 5 because a and b both are on same address

}
