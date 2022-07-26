package main

import (
    "fmt"
    "io/ioutil"
    "log"
)

func main() {

makeTree("designPattern")

}

func makeTree(dir string){
    files, err := ioutil.ReadDir(dir)

    if err != nil {

        log.Fatal(err)
    }

    for _, f := range files {

        if !f.IsDir() {
            fmt.Println("-----",f.Name())
            makeTree(f.Name())
        }else{
            fmt.Println(f.Name())
		}
    }
}