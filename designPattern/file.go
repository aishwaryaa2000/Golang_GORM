package main

import (
  "fmt"
  "log"
  "os"
  "bufio"
  "strings"

)

func main() {

  in := bufio.NewReader(os.Stdin)
  fmt.Print("Hey,please enter file path by having \\ in the path : ")
	filePath, _ := in.ReadString('\n')
	filePath = strings.TrimRight(filePath, "\r\n")
 // filePath = "C:\\Users\aishwarya.anand\\OneDrive - Forcepoint\\Desktop\\Swabhav_Techlabs"
  mainFolder := strings.Split(filePath, "\\")
  mainFolderName := mainFolder[len(mainFolder)-1]
  fmt.Println(mainFolderName)
  makeTree(filePath, 1)

}

func makeTree(dir string, level int) {
  f, err := os.Open(dir)
  if err != nil {
    log.Fatal(err)
  }

  files, err := f.Readdir(0)
  if err != nil {
    log.Fatal(err)
  }
  for _, f := range files {
    if f.IsDir() {
      dirName := dir + "/" + f.Name()
      for i := 0; i < level-1; i++ {
        fmt.Print("|\t")
      }
      fmt.Print("|------")
      fmt.Println(f.Name())

      makeTree(dirName, level+1)
    } else {

      for i := 0; i < level-1; i++ {
        fmt.Print("|\t")
      }
      fmt.Println("|-----", f.Name())
      fmt.Println("")

    }

  }
}