package main

import (
	"fmt"
	"reloaded/Functions"
	"os"
	"bufio"
)

func main() {
	// checking the args
	arg := os.Args[1:]
	if len(arg) != 2 {
		fmt.Println("ERROR: arg length")
		os.Exit(0)
	}
	// opening the file and checking weather it is empty or not
	file, err := os.Open(arg[0])
	if err != nil {
		fmt.Println("ERROR: File error")
		os.Exit(0)
	}
	// if len(os.ReadFile(arg[0])) <=0 {
	// 	fmt.Println("ERROR: File Empty")
	// 	os.Exit(0)
	// }
	defer file.Close()
	ResultFile, Rerr := os.Create(arg[1])
	if Rerr != nil  {
		fmt.Println("ERROR: Creating File")
		os.Exit(0)
	}
	//scanner to scan the file line by line
	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		if i == 1 {
			line := scanner.Text()
			line = reloaded.SpaceHandler(line)
			line = reloaded.Cap(line)
			line = reloaded.CapNum(line)
			line = reloaded.Low(line)
			line = reloaded.LowNum(line)
			line = reloaded.Up(line)
			line = reloaded.UpNum(line)
			line = reloaded.Vowels(line)
			line = reloaded.ConvertHextoDec(line)
			line = reloaded.ConvertBintoDec(line)
			line = reloaded.Punctuation(line)
			line = reloaded.SpaceHandler(line)
			line = reloaded.QuoteHandler(line)
			ResultFile.WriteString(line)
			i++
		} else {
			ResultFile.WriteString("\n")
			line := scanner.Text()
			line = reloaded.SpaceHandler(line)
			line = reloaded.Cap(line)
			line = reloaded.CapNum(line)
			line = reloaded.Low(line)
			line = reloaded.LowNum(line)
			line = reloaded.Up(line)
			line = reloaded.UpNum(line)
			line = reloaded.Vowels(line)
			line = reloaded.ConvertHextoDec(line)
			line = reloaded.ConvertBintoDec(line)
			line = reloaded.Punctuation(line)
			line = reloaded.SpaceHandler(line)
			line = reloaded.QuoteHandler(line)
			ResultFile.WriteString(line)
		}
		
		
		
	}
	ScanErr := scanner.Err()
	if ScanErr != nil {
		fmt.Println("ERROR: Scanner error")
		os.Exit(0)
	}

	
}
