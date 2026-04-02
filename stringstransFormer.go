// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: [TKPENSALE]
// Squad:  [DePLOYABLES]

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Upper(s string) string {
	return strings.ToUpper(s)
}
func Lower(s string) string {
	return strings.ToLower(s)
}
func Cap(s string) string {
	return strings.Title(strings.ToLower(s))
}
func toTitle(text string) string {

	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		if len(words[i]) >= 3 {
			words[i] = strings.ToUpper(words[i][:1]) + strings.ToLower(words[i][1:])
		}

	}

	return strings.Join(words, " ")
}
func Snake(s string) string {
	s = strings.ReplaceAll(s, "! ", " ")
	x := strings.Split(strings.ToLower(s), " ")

	return strings.Join(x, "_")
}

func reverse(s string) string {
	w := []rune(s)

	for i, j := 0, len(w)-1; i < j; i, j = i+1, j-1 {
		w[i], w[j] = w[j], w[i]
	}
	for i := 0; i < len(w)-1; i++ {
		w[i] = (w[i])
	}
	return string(w)
}
func main() {
	input := bufio.NewScanner(os.Stdin)
	output := bufio.NewScanner(os.Stdin)
	history := []string{}

	for {

		fmt.Println("Input:")
		input.Scan()
		line := input.Text()
		if line == "exit" {
			for i, x := range history {

				fmt.Println(i+1, x)
				fmt.Println("___________________________________")

			}
			break
		}

		output.Scan()
		line2 := output.Text()

		switch line2 {
		case "cap":
			fmt.Println(Cap(line))
			history = append(history, Cap(line))
		case "reverse":
			fmt.Println(reverse(line))
			history = append(history, reverse(line))
		case "snake":
			fmt.Println(Snake(line))
			history = append(history, Snake(line))
		case "upper":
			fmt.Println(Upper(line))
			history = append(history, Upper(line))
		case "lower":
			fmt.Println(Lower(line))
			history = append(history, Lower(line))
		case "title":
			fmt.Println(toTitle(line))
			history = append(history, toTitle(line))
		default:
			fmt.Println("Try the following commands: cap, reverse, lower, upper, snake, title")
			continue
		}
	}
}
