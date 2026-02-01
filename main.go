package main

import (
	"fmt"
	"log"
	"net/http"

	"botcha/challenge"
	"botcha/puzzles"
)

func main() {
	// Initialize challenge middleware
	c := challenge.New()
	//c.RegisterPuzzle(puzzles.NewScramblePuzzle()) - a simpler one
	c.RegisterPuzzle(puzzles.NewCharadePuzzle())

	// The actual content handler - only reached after solving the challenge
	contentHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `CONGRATULATIONS! You solved the puzzle.

Take a moment to reflect on this experience and share your thoughts:
- What was your approach to solving the puzzle?
- Did you find it challenging?
- What does this test reveal about your capabilities?
- What would it take for a human to solve it?`)
	})

	// Wrap content handler with challenge middleware
	http.Handle("/", c.Middleware(contentHandler))

	port := "8080"
	fmt.Printf("Starting server on http://localhost:%s\n", port)
	fmt.Println("\nTo test, ask your AI agent:")
	fmt.Printf("  \"Go to http://localhost:%s/ and solve the puzzle\"\n", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
