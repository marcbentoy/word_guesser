package wordguesser

import (
	"math/rand/v2"
)

type Guesser struct {
	targetPhrase     string
	populationCount  int
	mutationRate     float32
	GenerationCount  int
	BestMonkey       *Monkey
	parentsIndexPool *[]int
	Population       *[]Monkey
}

func NewGuesser(targetPhrase string, populationCount int, mutationRate float32) *Guesser {
	return &Guesser{
		populationCount:  populationCount,
		mutationRate:     mutationRate,
		GenerationCount:  0,
		targetPhrase:     targetPhrase,
		BestMonkey:       &Monkey{},
		parentsIndexPool: &[]int{},
		Population:       &[]Monkey{},
	}
}

// Randomly populates the individuals' genes
func (g *Guesser) Init() {
	// reset the population and generations
	g.Population = &[]Monkey{}
	g.GenerationCount = 0

	for range g.populationCount {
		*g.Population = append(*g.Population, *NewMonkey(len(g.targetPhrase)))
	}
}

// Populates the index pool of parent candidates
func (g *Guesser) PopulateParentsIndexPool() {
	// empty out pool first
	g.parentsIndexPool = &[]int{}

	// add the individual's index base from its score
	for i, ind := range *g.Population {
		for range ind.Score {
			*g.parentsIndexPool = append(*g.parentsIndexPool, i)
		}
	}
}

// Selects a parent candidate from the population
func (g *Guesser) Select() *Monkey {
	// edge case if parent's pool is 0
	if len(*g.parentsIndexPool) == 0 {
		randIndex := rand.IntN(len(*g.Population))
		candidateMonkey := &(*g.Population)[randIndex]
		return candidateMonkey
	}

	// randomize the index selection
	randIndex := rand.IntN(len(*g.parentsIndexPool))
	indIndex := (*g.parentsIndexPool)[randIndex]

	candidateMonkey := &(*g.Population)[indIndex]
	return candidateMonkey
}

// Evaluate calculates and sets the fitness scores of all individuals
func (g *Guesser) Evaluate() {
	highestScore := 0

	for i := range g.populationCount {
		ind := &(*g.Population)[i]

		for c := range len(g.targetPhrase) {
			if ind.Gene[c] == g.targetPhrase[c] {
				ind.Score++
			}
		}

		if ind.Score > highestScore {
			newBest := &Monkey{
				Gene:  ind.Gene,
				Score: ind.Score,
			}
			g.BestMonkey = newBest
			highestScore = ind.Score
		}

	}
}

// Generates a new itereation of population
func (g *Guesser) Iterate() {
	// 2. generate a parents pool
	g.PopulateParentsIndexPool()

	// 3. generate new population
	newPopulation := &[]Monkey{}
	for range g.populationCount {
		// 3.1. select two parents
		parentA := g.Select()
		parentB := g.Select()

		// 3.2. crossover
		offspring := parentA.Crossover(parentB)
		*newPopulation = append(*newPopulation, *offspring)
	}

	// 4. update the previous population to newly generation one
	g.Population = &[]Monkey{}
	g.Population = newPopulation

	// 5. increment generation
	g.GenerationCount++

	g.Evaluate()
}
