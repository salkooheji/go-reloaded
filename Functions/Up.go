package reloaded 

import (
	"fmt"
	// "regexp"
	// "strconv"
	"strings"
)

func Up(input string) string {
	words := strings.Fields(input)
	for i := 0; i < len(words); i++ {
		if strings.EqualFold(words[i], "(up)") {
			if i == 0 {
				fmt.Println("ERROR: '(up)' cannot be at the beginning of the input")
				continue
			}
			for j := i - 1; j >= 0; j-- {
				if IsAlphabetic(words[j]) {
					words[j] = strings.ToUpper(words[j])
					break
				}
			}
			words = append(words[:i], words[i+1:]...)
			i-- // Adjust index after removal
		}
	}
	return strings.Join(words, " ")
}
