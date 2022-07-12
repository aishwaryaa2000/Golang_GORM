package main

import(
	"battleShip/service"
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main(){
	in := bufio.NewReader(os.Stdin)

	var r uint8=11
	var c uint8 = 7
	board := service.Board(r,c)
	
	fmt.Print("Hey player,enter your name : ")
	name, _ := in.ReadString('\n')
	name = strings.TrimRight(name, "\r\n")

	player := service.Player(name)
	service.GameStart(board,player)

}
