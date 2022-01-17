package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data := []byte("Hello from go")
	err := ioutil.WriteFile("ourFile.txt",data,0644) //overwrites
	checkError(err)
	fmt.Println("Done writing")

	fptr, err := os.Create("ourFile1.txt")
	checkError(err)
	defer fptr.Close()
	bytesWritten, err := fptr.WriteString("Hello")
	checkError(err)
	fmt.Println("Bytes written are : ",bytesWritten)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error occurred : ", err)
	}
}