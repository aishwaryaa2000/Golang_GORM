package player

player1 :=''
player2 :=''

func PlayerInput(){
	fmt.Print("Hey Player1,choose X or O ")
	fmt.Scanln(&player1)
}