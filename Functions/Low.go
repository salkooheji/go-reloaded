package reloaded

import (
	"fmt"
	"strings"
)


func Low(input string) string {
	words := strings.Fields(input)
	for i := 0; i < len(words); i++ {
		if strings.EqualFold(words[i], "(low)") {
			if i == 0 {
				fmt.Println("ERROR: '(low)' cannot be at the beginning of the input")
				continue
			}
			// Changed line: lowercase the word before "(low)"
			words[i-1] = strings.ToLower(words[i-1])
			// Remove "(low)" as before
			words = append(words[:i], words[i+1:]...)
			i-- // Adjust index after removal
		}
	}
	return strings.Join(words, " ")
}
