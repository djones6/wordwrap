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
