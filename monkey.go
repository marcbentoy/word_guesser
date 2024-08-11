package wordguesser

import (
	"math/rand/v2"

	"github.com/marcbentoy/wordguesser/utils"
)

type Monkey struct {
	Gene  string
	Score int
}

func NewMonkey(length int) *Monkey {
	return &Monkey{
		Gene:  utils.RandomChars(length),
		Score: 0,
	}
}

// Randomly change the gene's value
func (m *Monkey) Mutate(mutationRate float32) {
	shouldMutate := rand.Float32()
	if shouldMutate <= mutationRate {
		// random index of the gene
		randIndex := rand.IntN(len(m.Gene))
		// random character to replace the previous gene character
		randChar := utils.RandomChars(1)

		// convert the gene to slice of runes
		indGene := []rune(m.Gene)
		indGene[randIndex] = rune(randChar[0])

		// update the mutated gene
		m.Gene = string(indGene)
	}
}

// Generates a new Monkey with a combined gene at a random point
func (m Monkey) Crossover(partner *Monkey) *Monkey {
	offspring := &Monkey{
		Gene:  "",
		Score: 0,
	}

	// point of intersection
	point := rand.IntN(len(m.Gene))

	offspring.Gene += m.Gene[:point]
	offspring.Gene += partner.Gene[point:]

	return offspring
}
