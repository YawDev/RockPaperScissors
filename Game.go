package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"os/exec"
)

var ComputerChoice string
var UserChoice string

func UserInput() {

	var userOption string
	var rounds int
	fmt.Println("Welcome to Rock Paper Scissors")

	for {
		fmt.Println("---Menu---")
		fmt.Println("1 - Start")
		fmt.Println("0 - Exit")
		fmt.Scan(&userOption)

		if userOption == "1" {
			Clear()
			fmt.Println("Game Started...")
			fmt.Println("How many rounds? ")
			fmt.Scan(&rounds)
			currentRound := 1
			for currentRound = 1; currentRound <= rounds; currentRound++ {
				fmt.Printf("Round %v:\n", currentRound)
				GameInProgress()
			}
		} else if userOption == "0" {
			break
		} else {
			fmt.Println("Enter a valid option")
		}

	}
	fmt.Println("Goodbye!")

}

func GameInProgress() {
	ComputerChoice = ComputerTurn()
	UserChoice = UserTurn()
	fmt.Println()
	fmt.Printf("You chose: %v\n", UserChoice)
	fmt.Printf("Computer chose: %v\n", ComputerChoice)
	fmt.Println()

	if ComputerHasWon() {
		fmt.Println("Computer Has Won!")
	} else if UserHasWon() {
		fmt.Println("You Won!")
	} else if IsTie() {
		fmt.Println("Game Tied!")
	}
	println()

}

func ComputerHasWon() bool {
	if ComputerChoice == "Paper" && UserChoice == "Rock" {
		return true
	}

	if ComputerChoice == "Scissors" && UserChoice == "Paper" {
		return true
	}

	if ComputerChoice == "Rock" && UserChoice == "Scissors" {
		return true
	}
	return false
}

func UserHasWon() bool {
	if ComputerChoice == "Paper" && UserChoice == "Scissors" {
		return true
	}

	if ComputerChoice == "Rock" && UserChoice == "Paper" {
		return true
	}

	if ComputerChoice == "Scissors" && UserChoice == "Rock" {
		return true
	}
	return false
}

func IsTie() bool {
	return ComputerChoice == UserChoice
}

func ComputerTurn() string {

	max := 3
	min := 1
	var computerChoice int = rand.IntN(max-min+1) + min

	if computerChoice == 1 {
		return "Rock"
	} else if computerChoice == 2 {
		return "Paper"
	}
	return "Scissors"
}

func UserTurn() string {

	for {
		fmt.Printf("Select (Rock (1), Paper (2), Or Scissors (3): ")
		var userChoice int
		fmt.Scan(&userChoice)

		if userChoice == 1 {
			return "Rock"
		} else if userChoice == 2 {
			return "Paper"
		} else if userChoice == 3 {
			return "Scissors"
		}
		fmt.Println("Not a valid option.")
	}

}

func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
