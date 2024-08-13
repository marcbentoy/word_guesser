package main

import (
	"fmt"

	wg "github.com/marcbentoy/wordguesser"
)

func main() {
	fmt.Println("Word Guesser made with Genetic Algorithms")

	targetPhrase := "vince gwapo kaayo to da max level"
	populationCount := 1000
	mutationRate := 0.01

	guesser := wg.NewGuesser(targetPhrase, populationCount, float32(mutationRate))

	fmt.Println("Initializing Population.")
	guesser.Init()
	fmt.Println("Population initialized.")
	guesser.Evaluate()

	for {
		fmt.Println()

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
