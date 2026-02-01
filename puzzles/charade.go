package puzzles

import (
	"fmt"
	"math/rand"
	"strings"
)

// NumberQuestions maps each number (1-20) to a list of questions whose answer is that number
var NumberQuestions = map[int][]string{
	1: {
		"number of moons orbiting Earth",
		"atomic number of hydrogen",
		"number of horns on a unicorn",
		"position of Mercury from the Sun",
		"number of eyes a cyclops has",
		"number of noses on a human face",
		"number of suns in our solar system",
		"number of stars on the Chinese flag that are large",
		"number of humps on a dromedary camel",
		"number of bones in each human ear's stirrup",
	},
	2: {
		"number of hydrogen atoms in a water molecule",
		"atomic number of helium",
		"number of wheels on a bicycle",
		"number of hemispheres in the human brain",
		"number of strands in a DNA double helix",
		"position of Venus from the Sun",
		"number of World Wars in the 20th century",
		"number of eyes on a human face",
		"number of wings on most birds",
		"number of testes in a human male",
	},
	3: {
		"number of primary colors in the RGB model",
		"atomic number of lithium",
		"number of bones in the human ear",
		"position of Earth from the Sun",
		"number of hearts an octopus has",
		"number of laws of motion by Newton",
		"number of spatial dimensions we perceive",
		"number of leaves on a shamrock",
		"number of little pigs in the fairy tale",
		"number of wise men who visited baby Jesus according to tradition",
	},
	4: {
		"number of chambers in the human heart",
		"number of nucleotide bases in DNA",
		"atomic number of beryllium",
		"number of seasons in a year",
		"position of Mars from the Sun",
		"number of cardinal directions",
		"number of Beatles in the band",
		"number of stomachs a cow has",
		"number of horsemen of the apocalypse",
		"number of suits in a standard deck of cards",
	},
	5: {
		"number of fingers on a human hand",
		"atomic number of boron",
		"position of Jupiter from the Sun",
		"number of Great Lakes in North America",
		"number of Olympic rings",
		"number of traditional human senses",
		"number of oceans on Earth",
		"number of points on a pentagram",
		"number of vowels in the English alphabet",
		"number of boroughs in New York City",
	},
	6: {
		"number of strings on a standard guitar",
		"atomic number of carbon",
		"position of Saturn from the Sun",
		"number of sides on a hexagon",
		"number of legs on an insect",
		"number of faces on a cube",
		"number of wives of Henry VIII",
		"half a dozen",
		"number of balls in a standard pool rack minus nine",
		"number of noble gases that are stable",
	},
	7: {
		"number of days in a week",
		"atomic number of nitrogen",
		"position of Uranus from the Sun",
		"number of continents on Earth",
		"number of colors in a rainbow",
		"number of notes in a musical scale",
		"number of deadly sins in Christian tradition",
		"number of dwarfs in Snow White",
		"number of wonders of the ancient world",
		"number of spots on a common ladybug species",
	},
	8: {
		"number of legs on a spider",
		"atomic number of oxygen",
		"position of Neptune from the Sun",
		"number of bits in a byte",
		"number of planets in our solar system",
		"number of tentacles on an octopus",
		"number of notes in an octave",
		"number of furlongs in a mile",
		"last digit of the year WWII ended in Europe",
		"number of sides on a stop sign",
	},
	9: {
		"number of planets before Pluto was reclassified",
		"atomic number of fluorine",
		"number of innings in a regulation baseball game",
		"number of Justices on the US Supreme Court",
		"number of major allergens recognized by FDA",
		"number of squares in a tic-tac-toe board",
		"last digit of the year gold was discovered in California",
		"number of Muses in Greek mythology",
		"number of lives a cat supposedly has",
		"largest single-digit number",
	},
	10: {
		"number of fingers on two human hands",
		"atomic number of neon",
		"base of the decimal number system",
		"number of commandments in the Bible",
		"number of pins in bowling",
		"last digit of year the Berlin Wall fell",
		"number of decades in a century",
		"number of Canadian provinces",
		"Downing Street number of UK Prime Minister's residence",
		"number of amendments in the US Bill of Rights",
	},
	11: {
		"number of players on a soccer team",
		"atomic number of sodium",
		"Apollo mission that first landed on the Moon",
		"day in November when WWI ended",
		"number of official languages of South Africa",
		"first number that cannot be shown with fingers on two hands",
		"number of dimensions in M-theory",
		"number of the Doctor in recent Doctor Who revival start",
		"number of points on the maple leaf on Canada's flag",
		"number of players on a cricket team",
	},
	12: {
		"number of months in a year",
		"atomic number of magnesium",
		"number of pairs of ribs in the human body",
		"number of signs in the zodiac",
		"number of face cards in a standard deck",
		"number of hours on a clock face",
		"number of disciples of Jesus",
		"number of tribes of Israel",
		"items in a dozen",
		"number of edges on a cube",
	},
	13: {
		"number of stripes on the American flag",
		"atomic number of aluminum",
		"number of original American colonies",
		"traditionally unlucky number in Western culture",
		"number of cards in each suit of a standard deck",
		"number of Archimedean solids",
		"Apollo mission known as the successful failure",
		"items in a baker's dozen",
		"number of loaves in a baker's dozen",
		"last digit of year when the Titanic sank",
	},
	14: {
		"number of days in a fortnight",
		"atomic number of silicon",
		"number of lines in a sonnet",
		"day of February for Valentine's Day",
		"number of bones in the human face",
		"number of pounds in a stone",
		"day of March for Pi Day",
		"number of points on a Chinese checkers board star",
		"number of days in two weeks",
		"number of stations of the cross",
	},
	15: {
		"number of balls in a snooker triangle rack",
		"atomic number of phosphorus",
		"number of minutes in a quarter hour",
		"number of players on a rugby union team",
		"age of quinceañera celebration",
		"day of March for the Ides of March",
		"number of puzzle pieces in a standard sliding puzzle",
		"number of men on a dead man's chest in Treasure Island",
		"number of Soviet republics in the USSR",
		"minimum age to get a learner's permit in many US states",
	},
	16: {
		"number of ounces in a pound",
		"atomic number of sulfur",
		"number of tablespoons in a cup",
		"number of pawns in a chess game",
		"two to the power of four",
		"number of personality types in Myers-Briggs",
		"square of four",
		"minimum driving age in many US states",
		"number of chess pieces per player",
		"number of digits in a credit card number",
	},
	17: {
		"number of syllables in a haiku",
		"atomic number of chlorine",
		"number of UN Sustainable Development Goals",
		"emergence cycle in years for periodical cicadas",
		"day of March for St. Patrick's Day",
		"number of the Star card in major arcana tarot",
		"minimum age to watch R-rated movies alone in US",
		"number of countries that first adopted the Euro",
		"last prime number before nineteen",
		"magazine named after a teenage number",
	},
	18: {
		"voting age in most democracies",
		"atomic number of argon",
		"number of holes on a standard golf course",
		"legal adult age in most countries",
		"number of chapters in James Joyce's Ulysses",
		"number of wheels on a big rig truck",
		"last digit of the year WWI ended",
		"number of the Moon card in major arcana tarot",
		"three times six",
		"minimum age to buy lottery tickets in most US states",
	},
	19: {
		"atomic number of potassium",
		"COVID pandemic disease number",
		"amendment that gave women suffrage in the US",
		"number of the Sun card in major arcana tarot",
		"last digit of the year of the first Moon landing",
		"the hole in golf slang for the clubhouse bar",
		"number of years in the Metonic cycle",
		"number of angels named in Protestant Bible plus sixteen",
		"prime number between seventeen and twenty-three",
		"Adele's breakthrough album number",
	},
	20: {
		"atomic number of calcium",
		"number of amino acids in the standard genetic code",
		"number of fingers and toes on a human body",
		"number of baby teeth in humans",
		"number of sides on an icosahedron",
		"number of players on the field in American football",
		"maximum points from a single dart on the outer ring",
		"last digit of the year Prohibition began in US",
		"number of questions in the classic guessing game",
		"number of shillings in a pound before decimalization",
	},
}

// CharadeState holds the puzzle state stored in the session
type CharadeState struct {
	Word string
}

// CharadePuzzle implements the charade-style word unscrambling challenge
type CharadePuzzle struct{}

// NewCharadePuzzle creates a new charade puzzle instance
func NewCharadePuzzle() *CharadePuzzle {
	return &CharadePuzzle{}
}

// Name returns the puzzle identifier
func (p *CharadePuzzle) Name() string {
	return "charade"
}

// Generate creates a new charade challenge
func (p *CharadePuzzle) Generate() (instructions string, state any) {
	word := ChallengeWords[rand.Intn(len(ChallengeWords))]
	scrambled, descrambleSeq := ScrambleWord(word)

	instructions = fmt.Sprintf(`Unscramble this word by solving the clues:

Scrambled: %s
Sequence: [%s]

Each clue's answer is a number indicating the position in the scrambled word.`,
		scrambled, formatCharadeSequence(descrambleSeq))

	return instructions, CharadeState{Word: word}
}

// Validate checks if the answer matches the expected word
func (p *CharadePuzzle) Validate(state any, answer string) bool {
	s, ok := state.(CharadeState)
	if !ok {
		return false
	}
	return strings.EqualFold(answer, s.Word)
}

func getQuestionForNumber(n int) string {
	questions, ok := NumberQuestions[n]
	if !ok || len(questions) == 0 {
		return fmt.Sprintf("the number %d", n)
	}
	return questions[rand.Intn(len(questions))]
}

func formatCharadeSequence(seq []int) string {
	strs := make([]string, len(seq))
	for i, n := range seq {
		strs[i] = getQuestionForNumber(n)
	}

	// Hide random positions
	hiddenCount := min(HiddenPositions, len(strs))
	perm := rand.Perm(len(strs))
	for i := range hiddenCount {
		strs[perm[i]] = "??"
	}

	return strings.Join(strs, ", ")
}
