package main

import (
	wg "github.com/marcbentoy/wordguesser"
	"github.com/marcbentoy/wordguesser/utils"
)

type Guesser struct {
	population *[]wg.Monkey
}

var targetPhrase = "vincegwapo"

var m1 = &wg.Monkey{
	Gene:  utils.RandomChars(len(targetPhrase)),
	Score: 0,
}
var m2 = &wg.Monkey{
	Gene:  utils.RandomChars(len(targetPhrase)),
	Score: 0,
}

var guesser = &Guesser{
	population: &[]wg.Monkey{
		*m1, *m2,
	},
}

func main() {
	wg.PrintPopulation(guesser.population)
	Evaluate()
	wg.PrintPopulation(guesser.population)
}

func Evaluate() {
	for i := range len(*guesser.population) {
		ind := &(*guesser.population)[i]

		for c := range len(targetPhrase) {
			if ind.Gene[c] == targetPhrase[c] {
				ind.Score++
			}
		}
	}
}
