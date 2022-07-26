package main

import (
  "fmt"
  "log"
  "os"
)

func main() {

  makeTree("C:\\Users\aishwarya.anand\\OneDrive - Forcepoint\\Desktop\\Swabhav_Techlabs", 1)

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