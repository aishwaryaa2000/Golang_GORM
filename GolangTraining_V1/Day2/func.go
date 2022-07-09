package main
import "fmt"

func main(){

    displayName()
    fmt.Println("End of main")
}

func displayName(){
    fmt.Println("Aishwarya Anand")
}


// If first letter of function name is small,
// then it is a private function and can't be accessed outside the package in which it is defined
// if first letter of the function is capital then it is public function