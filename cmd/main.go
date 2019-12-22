package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/rockettown1/learngo/rockPaperScissors/game"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(
		`
	So you think you can beat me at Rock. Paper. and Gopher? Pfft
	
	Very well, let's see what you've got.
	How many times would you like to play?
	`)
	scanner.Scan()
	rounds := scanner.Text()

	fmt.Println("	Ahhh " + rounds + " times ehhh? Well Good luck puny human, you're going to need it.")

	roundInt, _ := strconv.Atoi(rounds)

	for i := 0; i < roundInt; i++ {
		if i == 0 {
			fmt.Println("	Choose your weapon...")
		}
		fmt.Println("	Let's go again...")
		scanner.Scan()
		choice := scanner.Text()
		game.Logic(choice)
	}

}
