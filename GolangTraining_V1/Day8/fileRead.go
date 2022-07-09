package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//Reading all the content in the file

	content,err := ioutil.ReadFile("ourFile.txt")
	//content is slice of byte so UTF-8 values of characters
	checkError(err)
	fmt.Println("Content is  in UTF8 format of bytes: ",content)
	fmt.Println("Content when converted to string is : ",string(content))

	//Reading specific number of bytes from a file

	fptr,err := os.Open("ourFile.txt")
	defer fptr.Close()
	checkError(err)
	bucket := make([]byte,5) 
	//slice of bytes,only first 5 bytes will be stored in bucket
	bytesRead,err := fptr.Read(bucket)
	checkError(err)
	fmt.Println("Content of first 5 bytes is : ",string(bucket))
	fmt.Println("BytesRead is : ",bytesRead)
	
	// bucket = make([]byte,100) 
	// //slice of bytes,for first 100 bytes will be stored in bucket 
	// // so HELLO EVERYONE with 85 spaces will be stored
	// bytesRead,e := fptr.Read(bucket)
	// checkError(e)
	// fmt.Println("Content of first 5 bytes is : ",string(bucket[:bytesRead]))
	// fmt.Println("BytesRead is : ",bytesRead)


}

func checkError(err error){
	if err!=nil{
		fmt.Println("Error occurred : ",err)
	}
}