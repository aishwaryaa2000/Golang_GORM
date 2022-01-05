package main
import(
	"fmt"
	"strings"
)
func main(){

	str, strlen, _ := concatenateToUpper("Aishwarya","Anand")
	fmt.Println("Full name in upper case is : ",str)
	fmt.Println("And length of the string is : ",strlen)
}

func concatenateToUpper(firstName string, lastName string)(string,int,string){
  firstUpper := strings.ToUpper(firstName)
  lastUpper := strings.ToUpper(firstName)
  fullName := firstUpper + " " + lastUpper
  size := len(fullName)
  return fullName,size,"abc"

}

// func functionName(variableName dataType) (returnType1,returnType2,...) {
// if function returns 3 variables but if only 2 are used in the main func then use _ 

