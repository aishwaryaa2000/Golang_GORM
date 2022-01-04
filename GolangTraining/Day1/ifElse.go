package main

import "fmt"
import "time"

func main(){

  now:= time.Now().Hour()
  if now <= 12{
      fmt.Println("Good morning")
  } else if now <= 17{
      fmt.Println("Good Afternoon")
  } else if now <= 22{
      fmt.Println("Good Evening")
  } else {
     fmt.Println("Good night")
  }

  fmt.Println("Have a nice day")
}
