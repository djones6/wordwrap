package wordwrap

import (
	"strings"
)

// WordWrap wraps a string at the specified width
func WordWrap(theString string, width int) string {
	return wrap(theString, "", width)
}

func wrap(input string, output string, width int) string {
	if len(input) == 0 {
		return output
	}

	if len(input) <= width {
		return output + input
	}
	currentPiece := input[0 : width+1]

	lastWhitespaceIndex := strings.LastIndex(currentPiece, " ")
	if lastWhitespaceIndex > -1 {
		// freddy fox -> freddy\n
		output = output + currentPiece[0:lastWhitespaceIndex] + "\n"
		input = input[lastWhitespaceIndex+1:]
	} else {
		// bigbanana -> bigban\nana
		output = output + currentPiece[0:width] + "\n"
		input = input[width:]
	}

	return wrap(input, output, width)
}
