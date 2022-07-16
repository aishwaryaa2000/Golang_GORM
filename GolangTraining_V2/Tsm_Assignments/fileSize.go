package main

import (
    "fmt"
    "log"
    "os"
	"bufio"
	"strings"
)

type byteSize float64

const(
	_ =iota //0 value
	KB byteSize = 1<< (10*iota) //1024
	MB	
	GB
)

func main() {

	in := bufio.NewReader(os.Stdin)
	fmt.Print("Enter path of the file with extension : ")
	path, _ := in.ReadString('\n')
	path = strings.TrimRight(path, "\r\n")
	getFileSize(path)
	//DemoFileSize.txt
}

func getFileSize(path string){
	fileInfo, err := os.Stat(path)

    if err != nil {

        log.Fatal(err)
    }

    fileSize := byteSize(fileInfo.Size())

    fmt.Printf("The file size is %.2f B / %.2f KB / %.2f MB / %.2f GB \n", fileSize,fileSize/KB,fileSize/MB,fileSize/GB)
}