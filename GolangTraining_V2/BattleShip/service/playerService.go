package service

import (
	"battleShip/components/player"
	"os"
	"fmt"
	"bufio"
	"strings"
)

func MakePlayer() (*player.Player){
	in := bufio.NewReader(os.Stdin)
	fmt.Print("Hey player,enter your name : ")
	name, _ := in.ReadString('\n')
	name = strings.TrimRight(name, "\r\n")
	PlayerNow:= player.New(name)
	return PlayerNow
}

func takeInputFromPlayer(name string,rowSize,colSize uint8 )(uint8,uint8){
	var X,Y uint8
	begin1:
		fmt.Print(name," enter X cordinate where you wish to attack : ")
		_,err1 := fmt.Scan(&X)
		if err1!=nil || X>= rowSize{
			fmt.Println("Please enter an integer for board lesser than ",rowSize)
			goto begin1
		}
		
	begin2:
		fmt.Print(name," enter Y cordinate where you wish to attack : ")
		_,err2 := fmt.Scan(&Y)
		if err2!=nil || Y>= colSize{
			fmt.Println("Please enter an integer for board lesser than ",colSize)
			goto begin2
		}

	return X,Y
}
