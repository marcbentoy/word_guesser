package main

import (
	"fmt"

	wg "github.com/marcbentoy/wordguesser"
)

func main() {
	fmt.Println("Word Guesser made with Genetic Algorithms")

	targetPhrase := "I am who I am because of me that me is me"
	populationCount := 10000
	mutationRate := 0.01

	isAutomated := false

	guesser := wg.NewGuesser(targetPhrase, populationCount, float32(mutationRate))

	fmt.Println("Initializing Population.")
	guesser.Init()
	fmt.Println("Population initialized.")
	guesser.Evaluate()

	for {
		fmt.Println()
		if isAutomated {
			guesser.Iterate()
			fmt.Println("Generation: ", guesser.GenerationCount)
			fmt.Println("\nBest: ")
			wg.PrintMonkey(guesser.BestMonkey)

			if guesser.BestMonkey.Score != len(targetPhrase) {
				continue
			}

			fmt.Println("Best Monkey found: ", guesser.BestMonkey.Gene)
			break
		}

		fmt.Print("> ")
		command := ""
		fmt.Scanln(&command)

		if command == "r" {
			fmt.Println("Resetting generation..")
			fmt.Println("Initializing Population.")
			guesser.Init()
			fmt.Println("Population initialized.")
			continue
		}
		if command == "a" {
			if isAutomated {
				fmt.Println("Auto generation stopped.")
				isAutomated = false
				continue
			}
			isAutomated = true

			fmt.Println("Generating until maximum score is found..")
			continue
		}
		if command == "n" || command == "" {
			guesser.Iterate()
			fmt.Println("Generation: ", guesser.GenerationCount)
			fmt.Println("\nBest: ")
			wg.PrintMonkey(guesser.BestMonkey)
			continue
		}
		if command == "q" {
			fmt.Println("Quitting word guesser, bye!")
			break
		}
	}
}
