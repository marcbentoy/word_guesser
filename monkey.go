package wordguesser

import (
	ga "github.com/marcbentoy/wordguesser/ga"
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

// // Crossover implements ga.Individual.
// func (m Monkey) Crossover(*ga.Individual) *ga.Individual {
// 	panic("unimplemented")
// }

// // Mutate implements ga.Individual.
// func (m Monkey) Mutate() {
// 	panic("unimplemented")
// }

// Randomly change the gene's value
func (m *Monkey) Mutate() {

}

func (m *Monkey) Crossover(partner *ga.Individual) *ga.Individual {
	offspring := &Monkey{}

	var newInd ga.Individual = offspring

	return &newInd
}
