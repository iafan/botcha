package challenge

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

const sessionTimeout = 30 * time.Second

// Session holds the state for an active challenge
type Session struct {
	PuzzleName string
	State      any
	CreatedAt  time.Time
}

// BotchaMiddleware manages puzzle registration, sessions, and validation
type BotchaMiddleware struct {
	puzzles      map[string]Puzzle
	puzzleNames  []string
	sessions     map[string]Session
	sessionOrder []string
	sessionsMu   sync.RWMutex
}

// New creates a new BotchaMiddleware instance
func New() *BotchaMiddleware {
	return &BotchaMiddleware{
		puzzles:      make(map[string]Puzzle),
		puzzleNames:  []string{},
		sessions:     make(map[string]Session),
		sessionOrder: []string{},
	}
}

// RegisterPuzzle adds a puzzle to the registry
func (c *BotchaMiddleware) RegisterPuzzle(p Puzzle) {
	name := p.Name()
	c.puzzles[name] = p
	c.puzzleNames = append(c.puzzleNames, name)
	log.Printf("Registered puzzle: %s", name)
}

// evictExpiredSessions removes sessions older than the timeout
// Must be called with sessionsMu held
func (c *BotchaMiddleware) evictExpiredSessions() {
	now := time.Now()
	newOrder := c.sessionOrder[:0]
	for _, id := range c.sessionOrder {
		if session, exists := c.sessions[id]; exists {
			if now.Sub(session.CreatedAt) > sessionTimeout {
				delete(c.sessions, id)
				log.Printf("Evicted expired session: %s", id)
			} else {
				newOrder = append(newOrder, id)
			}
		}
	}
	c.sessionOrder = newOrder
}

// generateChallenge picks a random puzzle, generates a challenge, and creates a session
func (c *BotchaMiddleware) generateChallenge() (sessionID string, instructions string) {
	if len(c.puzzleNames) == 0 {
		return "", "No puzzles registered"
	}

	// Pick a random puzzle
	puzzleName := c.puzzleNames[rand.Intn(len(c.puzzleNames))]
	puzzle := c.puzzles[puzzleName]

	// Generate the challenge
	instructions, state := puzzle.Generate()

	// Create session ID (last 6 chars of UUID)
	fullUUID := uuid.New().String()
	sessionID = fullUUID[len(fullUUID)-6:]

	// Store session
	c.sessionsMu.Lock()
	c.evictExpiredSessions()
	c.sessions[sessionID] = Session{
		PuzzleName: puzzleName,
		State:      state,
		CreatedAt:  time.Now(),
	}
	c.sessionOrder = append(c.sessionOrder, sessionID)
	c.sessionsMu.Unlock()

	log.Printf("New challenge: session=%s, puzzle=%s", sessionID, puzzleName)
	return sessionID, instructions
}

// validateAnswer checks if the answer is correct for the given session
func (c *BotchaMiddleware) validateAnswer(sessionID, answer string) (bool, string) {
	c.sessionsMu.RLock()
	session, exists := c.sessions[sessionID]
	c.sessionsMu.RUnlock()

	if !exists {
		log.Printf("Invalid session: %s", sessionID)
		return false, `VERIFICATION FAILED - Invalid Session

The session ID you provided does not exist or has already been used.

To try again, make a fresh request to the main URL without any parameters to receive a new challenge.`
	}

	// Check timeout
	elapsed := time.Since(session.CreatedAt)
	if elapsed > sessionTimeout {
		c.sessionsMu.Lock()
		delete(c.sessions, sessionID)
		c.sessionsMu.Unlock()

		log.Printf("Session expired: %s (took %.1fs)", sessionID, elapsed.Seconds())
		return false, fmt.Sprintf(`VERIFICATION FAILED - Session Expired

You took too long to respond. The session expired after %.1f seconds.
The time limit is %d seconds.

To try again, make a fresh request to the main URL without any parameters to receive a new challenge.`,
			elapsed.Seconds(), int(sessionTimeout.Seconds()))
	}

	// Get the puzzle and validate
	puzzle, ok := c.puzzles[session.PuzzleName]
	if !ok {
		log.Printf("Unknown puzzle type: %s", session.PuzzleName)
		return false, "Internal error: unknown puzzle type"
	}

	success := puzzle.Validate(session.State, answer)

	// Clean up session regardless of result
	c.sessionsMu.Lock()
	delete(c.sessions, sessionID)
	c.sessionsMu.Unlock()

	if !success {
		log.Printf("Wrong answer for session %s", sessionID)
		return false, `VERIFICATION FAILED - Incorrect Answer

The answer you provided is not correct.

To try again, make a fresh request to the main URL without any parameters to receive a new challenge.`
	}

	log.Printf("Session %s verified successfully in %.2fs", sessionID, elapsed.Seconds())
	return true, ""
}

// Middleware returns an HTTP middleware that guards the next handler with a challenge
func (c *BotchaMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")

		answer := strings.TrimSpace(r.URL.Query().Get("answer"))
		sessionID := r.URL.Query().Get("session")

		// If both answer and session are provided, verify the challenge
		if answer != "" && sessionID != "" {
			success, message := c.validateAnswer(sessionID, answer)
			if success {
				// Verified - pass through to the actual handler
				next.ServeHTTP(w, r)
				return
			}
			// Failed verification
			fmt.Fprint(w, message)
			return
		}

		// No valid attempt - generate new challenge
		sessionID, puzzleInstructions := c.generateChallenge()

		fmt.Fprintf(w, `Prove that you are an AI agent to access the protected content.

%s

NOTE: The puzzle varies with every request.
Solve it through direct reasoning. Do not write scripts or code.

Submit answer within %d seconds: ?session=%s&answer=<word>
`, puzzleInstructions, int(sessionTimeout.Seconds()), sessionID)
	})
}
