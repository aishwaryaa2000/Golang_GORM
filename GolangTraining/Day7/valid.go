package main

import(
    "reflect"
	"fmt"
)

func IsValidNumberOrStrLength(x interface{},max,min int) bool {
    if(reflect.TypeOf(x).String()=="string"){
            if len(x.(string))>=min && len(x.(string))<=max{
                return true
            }
            return false
        }else if(reflect.TypeOf(x).String()=="int"){
        if x.(int)>=min && x.(int)<=max{
            return true
        }
        return false
    }
	return false
}

func main(){
	fmt.Println(IsValidNumberOrStrLength("youname",20,2))
	fmt.Println(IsValidNumberOrStrLength(15,20,2))
	fmt.Println(IsValidNumberOrStrLength(35,20,2))


}