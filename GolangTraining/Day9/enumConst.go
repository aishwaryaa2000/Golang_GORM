package main

import "fmt"
type byteSize float64

const(
	_ =iota //0 value
	KB byteSize = 1<< (10*iota) //1024
	MB
	GB
	TB
	PB
	EB
)

func main(){
	var fileSize byteSize = 5000000.0
	fmt.Printf("%.2fKB",fileSize/KB)
}