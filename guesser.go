package wordguesser

import (
	"math/rand/v2"

	"github.com/marcbentoy/wordguesser/ga"
)

type Guesser struct {
	targetPhrase     string
	populationCount  int
	mutationRate     float32
	generationCount  int
	parentsIndexPool []int
	population       *[]Monkey
}

func NewGuesser(targetPhrase string, populationCount int, mutationRate float32) *Guesser {
	return &Guesser{
		populationCount:  populationCount,
		mutationRate:     mutationRate,
		generationCount:  0,
		targetPhrase:     targetPhrase,
		parentsIndexPool: []int{},
		population:       &[]Monkey{},
	}
}

// Randomly populates the individuals' genes
func (g *Guesser) Init() {
	// reset the population and generations
	g.population = &[]Monkey{}
	g.generationCount = 0

	for range g.populationCount {
		*g.population = append(*g.population, *NewMonkey(len(g.targetPhrase)))
	}
}

// Selects a parent candidate from the population
func (g *Guesser) Select() *ga.Individual {
	randIndex := rand.IntN(len(g.parentsIndexPool))

	candidateMonkey := &(*g.population)[randIndex]
	var candidate ga.Individual = candidateMonkey
	return &candidate
}

// Populates the index pool of parent candidates
func (g *Guesser) PopulateParentsIndexPool() {
	// empty out pool first
	g.parentsIndexPool = []int{}

	// add the individual's index base from its score
	for i, ind := range *g.population {
		for range ind.Score {
			g.parentsIndexPool = append(g.parentsIndexPool, i)
		}
	}
}

// Evaluates the fitness scores of all individuals
func (g *Guesser) Evaluate() {
	for _, ind := range *g.population {
		for c := range len(g.targetPhrase) {
			if ind.Gene[c] == g.targetPhrase[c] {
				ind.Score++
			}
		}
	}
}

// Generates a new itereation of population
func (g *Guesser) Iterate() {
	// 1. evaluate the population
	g.Evaluate()

	// 2. generate a parents pool
	g.PopulateParentsIndexPool()

	// 3. generate new population
	newPopulation := &[]ga.Individual{}
	*newPopulation = append(*newPopulation, NewMonkey(10))

	// 3.1. select two parents
	parentA := g.Select()
	parentB := g.Select()

	// 3.2. crossover
	parentA.Crossover(parentB)
	// 3.3. check if newPopulation has populationCount
	// 4. increment generation
}
