package reloaded

import (
	
 	"strings"


)

func Vowels(input string) string {

	Vowels := []string {"a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H"}
	words := strings.Fields(input) 
	for i, word := range words {
		if word == "a" || word == "A" {
			for _, vowel := range Vowels {
				if string(words[i+1][0]) == vowel {
					if word == "a" {
						words[i] = "an"
					}
					if word == "A" {
						words[i] = "An"
					}
				}
			}
		}
	}
	return strings.Join(words, " ")
}