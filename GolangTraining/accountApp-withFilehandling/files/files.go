package files

import (
	"accountApp/user"
	"accountApp/account"
	"bufio"
	// "fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFromFile() {

	file, err := os.Open("accountApp.txt")
	defer file.Close()

	if err != nil {
		log.Fatalf("Failed to open")

	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	
	for _, each_ln := range text {

		eachWord := strings.Split(each_ln, ",")
		id,_ := strconv.Atoi(eachWord[0])
		name := eachWord[1]
		pass := eachWord[2]

		var userNew = user.New(name, pass,id)
		for i:=3;i<len(eachWord)-1;{
			accNo,_ := strconv.Atoi(eachWord[i])
			balance,_ := strconv.Atoi(eachWord[i+1])
			var newAcc = account.New(accNo,balance) 
			userNew.AddAccount(newAcc)
			i+=2
		}

		for i:=0;i<len(eachWord);i++{

		}
	}
}
