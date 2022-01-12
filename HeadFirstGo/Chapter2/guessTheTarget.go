package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1 
	fmt.Println("I have choosen a number between 1 to 100.\nCan you fuess it?")
	i:=10
	for ;i>0;i--{
		fmt.Printf("You have %d guesses left \n",i)
		fmt.Print("Make a guess : ")
		input,err := reader.ReadString('\n')
		if(err!=nil){
			log.Fatal(err)
		}
		input=strings.TrimSpace(input)
		guess,err := strconv.Atoi(input)
		if(err!=nil){
			log.Fatal(err)
		}
		if guess>target{
			fmt.Println("Oops! Your guess is high")
		}else if guess<target{
			fmt.Println("Oops! Your guess is low")
		}else{
			fmt.Println("Wohoo!You have guessed the number!")
			break
		}

	}

	if i==0{
		fmt.Println("Oops! You couldn't guess the number. It was : ",target)
	}
}