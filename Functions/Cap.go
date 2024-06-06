package reloaded

import (
	"fmt"
	"strings"
)

func Cap(input string) string {
	words := strings.Fields(input)
	for i := 0; i < len(words); i++ {
		if strings.EqualFold(words[i], "(cap)") {
			if i == 0 {
				fmt.Println("ERROR: '(cap)' cannot be at the beginning of the input")
				continue
			}
			// Changed line: capitalize the word before "(cap)"
			words[i-1] = Capitalize(words[i-1]) // Capitalize the first letter only
			// Remove "(cap)" as before
			words = append(words[:i], words[i+1:]...)
			i-- // Adjust index after removal
		}
	}
	return strings.Join(words, " ")
}
