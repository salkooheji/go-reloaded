package reloaded

import (
	"fmt"
	"strings"
	"strconv"
	"os"
)

func CapNum(input string) string {
	words := strings.Split(input, " ")
	for i := 0; i < len(words); i++ {
		if words[i] == "(cap," {
			if i == 0 {
				fmt.Println("ERROR: '(cap, <number>)' cannot be at the beginning of the input")
				os.Exit(0)
			}
			numC := strings.Trim(words[i+1], words[i+1][len(words[i+1])-1:])
			numCap, err := strconv.Atoi(numC)
			if err != nil {
				fmt.Println("ERROR: '(cap, <number>)' number has a problem")
				os.Exit(0)
			}
			if numCap > i || numCap < 0 {
				fmt.Println("ERROR: '(cap, <number>)' number out of range")
				os.Exit(0)
			}

			for j := 1; j <= numCap; j ++ {
				words[i-j] = Capitalize(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
		}	
	}
	return strings.Join(words, " ")
}
