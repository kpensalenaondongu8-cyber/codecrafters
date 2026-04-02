package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: go run . input.txt output.txt")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("read file successful")
	}

	words := (string(data))
	// processess all commmands such as caps, low, etc
	words = processCommands(words)
	// fix all qoutation errors
	words = fixQuotes(words)
	// fix all punctuations errors
	words = fixPunctuation(words)
	// process commands and write into ouput files
	words = fixGrammar(words)
	err = os.WriteFile(outputFile, []byte(words), 0644)
	if err == nil {
		fmt.Println("write file successful!")
	}

}
