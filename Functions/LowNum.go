package reloaded

import (
	"fmt"
	"strings"
	"strconv"
	"os"
)

func LowNum(input string) string {
	words := strings.Split(input, " ")
	for i := 0; i < len(words); i++ {
		if words[i] == "(low," {
			if i == 0 {
				fmt.Println("ERROR: '(low, <number>)' cannot be at the beginning of the input")
				os.Exit(0)
			}
			numL := strings.Trim(words[i+1], words[i+1][len(words[i+1])-1:])
			numLow, err := strconv.Atoi(numL)
			if err != nil {
				fmt.Println("ERROR: '(low, <number>)' number has a problem")
				os.Exit(0)
			}
			if numLow > i || numLow < 0 {
				fmt.Println("ERROR: '(low, <number>)' number out of range")
				os.Exit(0)
			}

			for j := 1; j <= numLow; j ++ {
				words[i-j] = strings.ToLower(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
		}	
	}
	return strings.Join(words, " ")
}
