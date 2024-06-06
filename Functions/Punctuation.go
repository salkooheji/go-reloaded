package reloaded

import (
    "strings"
)


func Punctuation(input string) string {
	input = strings.ReplaceAll(input, " ?", "?")
	input = strings.ReplaceAll(input, " .", ".")
	input = strings.ReplaceAll(input, " ,", ",")
	input = strings.ReplaceAll(input, " !", "!")
	input = strings.ReplaceAll(input, " :", ":")
	input = strings.ReplaceAll(input, " ;", ";")
	return input
}
