package main
import (
	"fmt"
	"reflect"
)
func main(){
	fmt.Println(reflect.TypeOf("Hello"))
    fmt.Println(reflect.TypeOf(4.4))
    fmt.Println(reflect.TypeOf(100))
    fmt.Println(reflect.TypeOf(true))
}