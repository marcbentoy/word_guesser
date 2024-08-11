package ga

type Individual interface {
	// Mutates the gene of an individual
	Mutate()

	// Creates a new offspring with the specified partner
	Crossover(*Individual) *Individual
}
