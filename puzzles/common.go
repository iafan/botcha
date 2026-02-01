package puzzles

import (
	"math/rand"
	"strings"
)

// Shared constants for word scramble puzzles
const (
	HiddenPositions = 2
	ExtraLetters    = 5
)

// ChallengeWords is the shared list of words used by scramble-based puzzles
var ChallengeWords = []string{
	"constantinople",
	"flabbergasted",
	"discombobulated",
	"onomatopoeia",
	"metamorphosis",
	"philanthropist",
	"extraterrestrial",
	"bioluminescence",
	"circumnavigation",
	"perpendicular",
	"gubernatorial",
	"idiosyncratic",
	"quintessential",
	"czechoslovakia",
	"enthusiastically",
	"incomprehensible",
	"procrastination",
	"anthropomorphic",
	"compartmentalize",
	"differentiation",
	"electromagnetic",
	"fundamentalism",
	"hallucination",
	"jurisprudence",
	"knowledgeable",
	"mediterranean",
	"neuroscientist",
	"organizational",
	"photosynthesis",
	"questionnaire",
	"reconnaissance",
	"rehabilitation",
	"sophisticated",
	"telecommunications",
	"underestimating",
	"ventriloquist",
	"weightlessness",
	"autobiographical",
	"disproportionate",
	"counterproductive",
	"microorganisms",
	"prestidigitation",
}

// ScrambleWord shuffles a word, adds extra random letters, and returns the
// scrambled version along with the descramble sequence (1-indexed positions)
func ScrambleWord(word string) (string, []int) {
	runes := []rune(strings.ToLower(word))
	n := len(runes)

	// Track where each original position ends up
	indices := make([]int, n)
	for i := range indices {
		indices[i] = i
	}

	// Shuffle both runes and indices together
	// Keep reshuffling until first 3 original positions aren't in first 3 scrambled positions
	for {
		rand.Shuffle(n, func(i, j int) {
			runes[i], runes[j] = runes[j], runes[i]
			indices[i], indices[j] = indices[j], indices[i]
		})

		// Check constraint: none of indices[0], indices[1], indices[2] should be 0, 1, or 2
		valid := true
		for scrambledPos := 0; scrambledPos < 3 && scrambledPos < n; scrambledPos++ {
			if indices[scrambledPos] < 3 {
				valid = false
				break
			}
		}
		if valid {
			break
		}
	}

	// Build descramble sequence: for each position in original word,
	// which position in scrambled word has that letter?
	// indices[scrambledPos] = originalPos, so we invert it
	descrambleSeq := make([]int, n)
	for scrambledPos, originalPos := range indices {
		descrambleSeq[originalPos] = scrambledPos + 1 // 1-indexed
	}

	// Add extra random letters to increase difficulty
	for range ExtraLetters {
		letter := rune('a' + rand.Intn(26))
		pos := rand.Intn(len(runes) + 1)

		runes = append(runes[:pos], append([]rune{letter}, runes[pos:]...)...)

		for i, p := range descrambleSeq {
			if p >= pos+1 {
				descrambleSeq[i] = p + 1
			}
		}
	}

	return string(runes), descrambleSeq
}

