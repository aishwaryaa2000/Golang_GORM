package main

import "fmt"

func main() {
	myMap := map[int]string{
		1: "a",
		2: "b",
		3: "c", //notice this , in the end.
	}

	newMap := checkMap(myMap)
	fmt.Println("My map : ",myMap)
	fmt.Println("My newmap : ",newMap)
	
	// map are reference type,any changes in myMap will be reflected in newMap
    myMap[1]="yo"
    fmt.Println("My map : ",myMap)
	fmt.Println("My newmap : ",newMap)

}

func checkMap(myMap map[int]string) map[int]string {
	myMap[1] = "Changed via func" 
	//myMap changed here, will get reflected in myMap in main() as maps are referenced,just like slices
	return myMap
}

