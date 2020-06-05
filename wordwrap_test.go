package wordwrap

// Your task is to write a function that takes two arguments, a string and an integer width.
// The function returns the string, but with line breaks inserted at just the right places to make sure that no line is longer than the column number.
// You try to break lines at word boundaries.
// Like a word processor, break the line by replacing the last space in a line with a newline.
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test that an empty string is returned unmodified.
func TestEmptyString(t *testing.T) {
	emptyString := ""
	assert.Equal(t, emptyString, WordWrap(emptyString, 0))
}

// Test that a string shorter than width is returned unmodified.
func TestStringShorterThanWidth(t *testing.T) {
	input := "banana"
	assert.Equal(t, input, WordWrap(input, 7))
}

// Test that a single word that is longer than the maximum width is split onto multiple lines.
func TestSingleWordLongerThanWidth(t *testing.T) {
	input := "bigbanana"
	expected := "bigbana\nna"
	assert.Equal(t, expected, WordWrap(input, 7))
}

func TestWrapTwoWords(t *testing.T) {
	input := "freddy fox"
	expected := "freddy\nfox"
	assert.Equal(t, expected, WordWrap(input, 7))
}

func TestWrapThreeWordsOnSeparateLines(t *testing.T) {
	input := "freddy farty fox"
	expected := "freddy\nfarty\nfox"
	assert.Equal(t, expected, WordWrap(input, 6))
}

func TestWrapOnWordBoundary(t *testing.T) {
	input := "freddy fox"
	expected := "freddy\nfox"
	assert.Equal(t, expected, WordWrap(input, 6))
}

func TestWordSplitAcrossMultiLines(t *testing.T) {
	input := "Supercalifragilisticexpialidocious"
	expected := "Supercalif\nragilistic\nexpialidoc\nious"
	assert.Equal(t, expected, WordWrap(input, 10))
}

func TestSuperCaliFragiWhatnot(t *testing.T) {
	input := "It's supercalifragilisticexpialidocious Even though the sound of it Is something quite atrocious"
	expected := "It's\nsupercalifragilistic\nexpialidocious Even\nthough the sound of\nit Is something\nquite atrocious"
	assert.Equal(t, expected, WordWrap(input, 20))
}

func TestEssay(t *testing.T) {
	input := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
	expected := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor\n" +
		"incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud\n" +
		"exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute\n" +
		"irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla\n" +
		"pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia\n" +
		"deserunt mollit anim id est laborum."
	assert.Equal(t, expected, WordWrap(input, 84))
}

func TestTrailingSpace(t *testing.T) {
	input := "Vanilla   "
	expected := "Vanilla   "
	assert.Equal(t, expected, WordWrap(input, 15))
}

func TestLeadingSpace(t *testing.T) {
	input := "   Vanilla"
	expected := "   Vanilla"
	assert.Equal(t, expected, WordWrap(input, 15))
}
