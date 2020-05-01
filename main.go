package wordwrap

import (
	"fmt"
	"net/http"
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

	for index, char := range theString {
		if char == ' ' {

		}
	}

	return theString[0:width] + "\n" + theString[width:len(theString)]
}
