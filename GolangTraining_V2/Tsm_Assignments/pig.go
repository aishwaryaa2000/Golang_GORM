/* 
Using double for loop.One in main and one in number

For double for loop and use of pointers, refer pig.go
*/

package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func turnScore(iTurn uint64, totalScore *uint64) bool {

	fmt.Println("\nThis is your turn ", iTurn)
	rollOrHold := "r"
	var (
		currentTurnScore, diceRoll uint64 = 0, 0
	)
	for {

		fmt.Print("Enter 'r' to roll again, 'h' to hold : ")
		fmt.Scan(&rollOrHold)
		if rollOrHold == "h" {
			*totalScore = *totalScore + currentTurnScore
			fmt.Println("You have holded.Your turn score is ", currentTurnScore, " and your new total score is", *totalScore)
			return false

		}

		if rollOrHold != "r" {
			//When user has entered h then the above if will be executed
			//If user has entered something which is not r then he has entered an invalid input
			fmt.Println("You have entered incorrect character.")
			continue
		}

		n, _ := rand.Int(rand.Reader, big.NewInt(5))
		diceRoll = uint64(n.Int64() + 1) //number between 1 to 6 is generated
		fmt.Println("You rolled : ", diceRoll)
		if diceRoll == 1 {
			fmt.Println("Oops!Turn over. No score for this turn")
			currentTurnScore = 0
			break
		}

		currentTurnScore = currentTurnScore + diceRoll
		fmt.Println("\nYour turn score is ",currentTurnScore," and your total score is", *totalScore,"\nIf you hold then your total score will be ",currentTurnScore+ *totalScore)

		if *totalScore+currentTurnScore >= uint64(20) {
			fmt.Println("\nYou Win,you finished in ", iTurn, "turns!\nNew total score considering the current total score and turn score is : ", *totalScore+currentTurnScore)
			return true
		}

	}

	return false
}

func main() {

	var (
		totalScore, iTurn uint64 = 0, 0
	)

	fmt.Println("Let's play Pig Dice game")
	for {
		iTurn = iTurn + 1
		okGameFinish := turnScore(iTurn, &totalScore)
		if okGameFinish {
			return
		}

	}
}
	 