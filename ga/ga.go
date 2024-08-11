package ga

type GA interface {
	// Initializes a population which is a slice of individuals.
	// This should also sets the first generation of that population.
	Init(*[]Individual)

	// Selects a new candidate for a parent of the given population.
	Select(*[]Individual) *Individual

	//
	Evaluate()
}
