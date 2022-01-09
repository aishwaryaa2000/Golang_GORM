package main
import(
	"fmt"
	"reflect"
)

func main(){
	var myArray = [5]string{"Hello","from","go","lang"}
	fmt.Println(reflect.TypeOf(myArray))  
	fmt.Println("By default value of myArray[4] : ",myArray[4])  
	// prints [4]string i.e [sizeofarray]string and if it was an integer array then [size]int

	fmt.Println("Inside main() the value of myArray is : ",myArray)
	case2(&myArray)

	for index,value := range myArray{
		fmt.Println("Index is : ",index,"and value is : ",value)
	}


}

func case1(arr [5]string){
	arr[0]= "Hey"
	fmt.Println("Inside case1() the value of arr is : ",arr)
	// call by value so changes in case1() is not reflected in main()
}

func case2(arrptr *[5]string){
	fmt.Println(arrptr[0])
	fmt.Println((*arrptr)[0])
}
