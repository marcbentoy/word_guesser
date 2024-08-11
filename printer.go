package wordguesser

import (
	"fmt"
)

func PrintMonkey(ind *Monkey) {
	fmt.Println(ind.Gene, "\t", ind.Score)
}

func PrintPopulation(population *[]Monkey) {
	for _, ind := range *population {
		PrintMonkey(&ind)
	}
}
