package reloaded

import (
	"strings"
)

func QuoteHandler(input string) string {
	counter := 0
	words := strings.Fields(input)
	for i, word := range words {
		if word == "'" {
			counter ++
		}
		if word == "'" && (counter%2 != 0){
			if i+2 > len(words) {
				break
			}
			words[i+1] = word + words[i+1]
			words = append(words[:i], words[i+1:]...)
		}
		if word == "'" && (counter%2 == 0) {
			words[i-1] = words[i-1] + word
			words = append(words[:i], words[i+1:]...)
		}
	}
	input = strings.Join(words, " ")
	return input
}