package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	if fileSize/MB > 50{
		log.Fatal("Max file size can only be till 50MB and not greater than that")
	}

    fmt.Printf("The file size is %.2f B / %.2f KB / %.2f MB / %.2f GB \n", fileSize,fileSize/KB,fileSize/MB,fileSize/GB)
}