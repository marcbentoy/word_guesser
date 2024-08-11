package main

import (
	"fmt"

	wg "github.com/marcbentoy/wordguesser"
)

func main() {
	fmt.Println("Word Guesser made with Genetic Algorithms")

	targetPhrase := "vince gwapo"
	populationCount := 1000
	mutationRate := 0.2

	guesser := wg.NewGuesser(targetPhrase, populationCount, float32(mutationRate))

	fmt.Println("Initializing Population.")
	guesser.Init()
	fmt.Println("Population initialized.")
	fmt.Println("Initial Population:")
	guesser.Evaluate()

	for {
		fmt.Println()

		fmt.Print("> ")
		command := ""
		fmt.Scanln(&command)

		if command == "n" || command == "" {
			guesser.Iterate()
			fmt.Println("Generation: ", guesser.GenerationCount)
			fmt.Println("Population: ")
			wg.PrintPopulation(guesser.Population)
			fmt.Println("\nBest: ")
			wg.PrintMonkey(guesser.BestMonkey)
		}
		if command == "quit" {
			fmt.Println("Quitting word guesser, bye!")
			break
		}
	}
}
