package reloaded

import (
	"strings"
	"strconv"
	"fmt"
	"os"
	"log"
)

func ConvertHextoDec(input string) string {
	words := strings.Fields(input)
	for i, word := range words {
		if word == "(hex)" {
			if i == 0 {
				fmt.Println("ERROR: '(hex)' cannot be at the beginning of the input")
				os.Exit(0)			
			}
			temp, err := strconv.ParseInt(words[i-1], 16, 64)
			if err!= nil{
				log.Fatal(err)
			}
			words[i-1] = strconv.FormatInt(temp, 10)
			words = append(words[:i], words[i+1:]...)
		}
	}
	return strings.Join(words, " ")
}
