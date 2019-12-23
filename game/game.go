package game

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//Logic for the Rock Paper Scissors game, generates a random choice for the opponents then determines a winner
func Logic(choice string) int {
	// generate a random number between 0 and 2
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(3)

	weapons := [3]string{"rock", "paper", "gopher"}
	opponentChoice := weapons[n]

	c := strings.ToLower(choice)

	if c == opponentChoice {
		fmt.Println("	We both picked the same. Pfft You must have copied me.")
	} else if c == "rock" {
		if opponentChoice == "gopher" {
			fmt.Println("	You picked " + c)
			fmt.Println("	I picked " + opponentChoice)
			fmt.Println("	WHAT?! That cannot be? You must have cheated")
			return 1
		} else if opponentChoice == "paper" {
			fmt.Println("	You picked " + c)
			fmt.Println("	I picked " + opponentChoice)
			fmt.Println("	HAHA! You'll never beat me.")
			return 0
		}
	} else if c == "gopher" {
		if opponentChoice == "rock" {
			fmt.Println("	You picked " + c)
			fmt.Println("	I picked " + opponentChoice)
			fmt.Println("	HAHA! You'll never beat me.")
			return 0
		} else if opponentChoice == "paper" {
			fmt.Println("	You picked " + c)
			fmt.Println("	I picked " + opponentChoice)
			fmt.Println("	WHAT?! That cannot be? You must have cheated")
			return 1
		}
	} else if c == "paper" {
		if opponentChoice == "rock" {
			fmt.Println("	You picked " + c)
			fmt.Println("	I picked " + opponentChoice)
			fmt.Println("	WHAT?! That cannot be? You must have cheated")
			return 1
		} else if opponentChoice == "gopher" {
			fmt.Println("	You picked " + c)
			fmt.Println("	I picked " + opponentChoice)
			fmt.Println("	HAHA! You'll never beat me.")
			return 0
		}
	} else {
		fmt.Println("	Hahah, you didn't choose a proper weapon. Are you scared?")
	}
	return 2
}
