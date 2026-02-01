package challenge

// Puzzle defines the interface that all challenge puzzles must implement
type Puzzle interface {
	// Name returns the unique identifier for this puzzle type
	Name() string

	// Generate creates a new challenge and returns:
	// - instructions: the text to show the user/agent
	// - state: puzzle-specific data to store in the session for validation
	Generate() (instructions string, state any)

	// Validate checks if the provided answer is correct for the given state
	Validate(state any, answer string) bool
}
