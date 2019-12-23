package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/rockettown1/learngo/rockPaperScissors/game"
)

func main() {
	//initialise scanner to read command line user input
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(
		`
	So you think you can beat me at Rock. Paper. and Gopher? Pfft
	
	Very well, let's see what you've got.
	How many times would you like to play?
	`)
	scanner.Scan()
	rounds := scanner.Text()
	roundInt, err := strconv.Atoi(rounds)

	//if the player entered something which wasn't a number, give them a second chance to enter correctly
	if err != nil {
		fmt.Println("	You must select a valid number of times to play...")
		scanner.Scan()
		rounds = scanner.Text()
		rountInt2, err := strconv.Atoi(rounds)
		if err != nil {
			fmt.Println("	Hmm too scared to play ehh? Well, until next time.")
			return
		} else {
			fmt.Println("	Ahhh " + rounds + " times ehhh? Well Good luck puny human, you're going to need it.")
			roundInt = rountInt2
		}
	} else {
		fmt.Println("	Ahhh " + rounds + " times ehhh? Well Good luck puny human, you're going to need it.")
	}

	//initialise scores to start the game
	scoreHuman := 0
	scoreOpp := 0

	//loop for each round
	for i := 0; i < roundInt; i++ {
		if i == 0 {
			fmt.Println("	Choose your weapon...")
		} else {
			fmt.Println("	Let's go again...")
		}

		scanner.Scan()
		choice := scanner.Text()
		//each choice compared to the opponents randomly generated one. returns 1 if human scores or 0 if opponent scores
		val := game.Logic(choice)

		if val == 1 {
			scoreHuman++
		} else if val == 0 {
			scoreOpp++
		}

		fmt.Println("	Your score: " + strconv.Itoa(scoreHuman))
		fmt.Println("	My score: " + strconv.Itoa(scoreOpp))
		fmt.Println()
	}

	if scoreHuman == scoreOpp {
		fmt.Println("	Hmm a draw? You are very lucky human. Next time you won't be!")
	} else if scoreHuman > scoreOpp {
		fmt.Println("	%&!? NO. HOW. This is impossible!! Next time we play I will crush your soul.")
	} else {
		fmt.Println("	HAHAHA. I laugh at your rock, paper, gopher skills.")
	}

}
