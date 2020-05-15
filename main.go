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

	firstBit := strings.TrimSpace(theString[0:width])
	lastBit := strings.TrimSpace(theString[width:len(theString)])
	return firstBit + "\n" + lastBit
}
