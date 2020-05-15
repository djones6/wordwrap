package wordwrap

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}

// WordWrap wraps a string at the specified width
func WordWrap(theString string, width int) string {
	if len(theString) <= width {
		return theString
	}

	// Consider string[index:index+width+1] and look backward for whitespace
	// if whitespace found split on it
	// otherwise split at width
	// index = whatever you split on
	// carry on

	currentPiece := theString[0 : width+1]
	resultString := wrap(currentPiece, width)

	firstBit := strings.TrimSpace(theString[0:width])
	lastBit := strings.TrimSpace(theString[width:len(theString)])
	return firstBit + "\n" + lastBit
}

// global resultString
// func wrap (theString string, currentIndex int, width int) {
//   if FINISHED return
//   currentPiece := theString[currentIndex : width+1]
//   resultString += Work out substring + '\n'
//   return wrap(whitespaceindex:nextIndexAtSizeWidth)
// }

func wrap(currentPiece string, width int) string {
	lastWhitespaceIndex := strings.LastIndex(currentPiece, " ")
	if lastWhitespaceIndex > -1 {
		return currentPiece[0:lastWhitespaceIndex] + "\n"
	}

	return currentPiece[0:width]
}
