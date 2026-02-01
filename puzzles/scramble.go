package puzzles

import (
	"fmt"
	"math/rand"
	"strings"
)

// ScrambleState holds the puzzle state stored in the session
type ScrambleState struct {
	Word string
}

// ScramblePuzzle implements the word unscrambling challenge
type ScramblePuzzle struct{}

// NewScramblePuzzle creates a new scramble puzzle instance
func NewScramblePuzzle() *ScramblePuzzle {
	return &ScramblePuzzle{}
}

// Name returns the puzzle identifier
func (p *ScramblePuzzle) Name() string {
	return "scramble"
}

// Generate creates a new scramble challenge
func (p *ScramblePuzzle) Generate() (instructions string, state any) {
	word := ChallengeWords[rand.Intn(len(ChallengeWords))]
	scrambled, descrambleSeq := ScrambleWord(word)

	instructions = fmt.Sprintf(`Unscramble this word:

Scrambled: %s
Sequence: [%s]`,
		scrambled, formatSequence(descrambleSeq))

	return instructions, ScrambleState{Word: word}
}

// Validate checks if the answer matches the expected word
func (p *ScramblePuzzle) Validate(state any, answer string) bool {
	s, ok := state.(ScrambleState)
	if !ok {
		return false
	}
	return strings.EqualFold(answer, s.Word)
}

var numberWords = []string{
	"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen",
	"seventeen", "eighteen", "nineteen", "twenty",
}

func numberToWord(n int) string {
	if n >= len(numberWords) {
		return fmt.Sprintf("%d", n)
	}
	word := numberWords[n]
	if len(word) > 4 {
		// Remove a random letter to prevent simple scripting
		removeIdx := rand.Intn(len(word))
		word = word[:removeIdx] + word[removeIdx+1:]
	}
	return word
}

func formatSequence(seq []int) string {
	strs := make([]string, len(seq))
	for i, n := range seq {
		strs[i] = numberToWord(n)
	}

	// Hide random positions
	hiddenCount := min(HiddenPositions, len(strs))
	perm := rand.Perm(len(strs))
	for i := range hiddenCount {
		strs[perm[i]] = "--"
	}

	return strings.Join(strs, ", ")
}
