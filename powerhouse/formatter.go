package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func formatter() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("WELCOME! THIS IS A STRING FORMATTER CLI")
	fmt.Println()

initial:
	fmt.Println("Our Operators are: \n-upper <text> \n-lower <text>\n-cap <text>\n-title <text>\n-snake <text>\n-reverse <text>\n-exit ")
	fmt.Println()
	fmt.Println("choose an operator!")
	scanner.Scan()

	input := scanner.Text()
	input = strings.TrimSpace(input)

	parts := strings.SplitN(input, " ", 2)
	operation := parts[0]
	operation = strings.ToLower(string(operation))

	if operation == "exit" {
		fmt.Println("Shutting down please wait...")
		fmt.Println("see you soon buddy!")
		powerHouse()
	}

	var text string
	if len(parts) == 2 {
		text = parts[1]
		words := strings.Fields(text)
		text = strings.Join(words, " ")
	}

	if len(parts) != 2 {
		fmt.Println("✗ No text provided. Usage: "+"operation", "+", "<text>")
		fmt.Println()
		goto initial
	}

	switch operation {
	case "cap":
		fmt.Println(titleCase(text))
		fmt.Println()
		history = append(history, operation+" "+text)
		goto initial
	case "upper":
		fmt.Println(ToUpperCase(text))
		fmt.Println()
		history = append(history, operation+" "+text)
		goto initial
	case "lower":
		fmt.Println(ToLowerCase(text))
		fmt.Println()
		goto initial
	case "title":
		fmt.Println(capsFirstLetter(text))
		fmt.Println()
		history = append(history, operation+" "+text)
		goto initial
	case "snake":
		fmt.Println(snakeCase(text))
		fmt.Println()
		history = append(history, operation+" "+text)
		goto initial
	case "reverse":
		fmt.Println(reverseSentence(text))
		fmt.Println()
		history = append(history, operation+" "+text)
		goto initial
	default:
		fmt.Println(`✗ Unknown command: ` + fmt.Sprintf("%q", operation))
		goto initial
	}

}
