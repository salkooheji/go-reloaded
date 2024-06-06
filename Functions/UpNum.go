package reloaded

import (
	"fmt"
	"strings"
	"strconv"
	"os"
)

func UpNum(input string) string {
	words := strings.Split(input, " ")
	for i := 0; i < len(words); i++ {
		if words[i] == "(up," {
			if i == 0 {
				fmt.Println("ERROR: '(up, <number>)' cannot be at the beginning of the input")
				os.Exit(0)
			}
			numU := strings.Trim(words[i+1], words[i+1][len(words[i+1])-1:])
			numUp, err := strconv.Atoi(numU)
			if err != nil {
				fmt.Println("ERROR: '(up, <number>)' number has a problem")
				os.Exit(0)
			}
			if numUp > i || numUp < 0 {
				fmt.Println("ERROR: '(up, <number>)' number out of range")
				os.Exit(0)
			}

			for j := 1; j <= numUp; j ++ {
				words[i-j] = strings.ToUpper(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
		}	
	}
	return strings.Join(words, " ")
}
