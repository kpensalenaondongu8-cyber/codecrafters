package main

import (
	"fmt"
	"os"
)

var history []string

func powerHouse() {

	var (
		tokens int
		help   string
	)

	help = "calc → the calculator\nbase → the base converter\nstr → the string transformer\nhelp → shows all commands\nhistory → shows last 10 inputs\nexit→ shuts down the console"
	fmt.Println("Welcome to command && Control Console")
initial:
	fmt.Println("Option 1 calc | option 2 base_converter | Option 3 str formatter | option 4 help | option 5 history | option 6 exit")

	for {
		n, err := fmt.Scan(&tokens)
		if err != nil && n != 1 {
			fmt.Println(err)
			var discard string
			fmt.Scan(&discard)
			goto initial
		}

		if tokens < 0 && tokens > 6 {
			fmt.Println("Invalide COmmand token")
			goto initial
		}

		if tokens == 1 {
			calculator()
		}
		if tokens == 2 {
			baseConverter()
		}
		if tokens == 3 {
			formatter()
		}
		if tokens == 4 {
			fmt.Println(help)
			goto initial
		}

		if tokens == 5 {
			fmt.Println("History Log!!")
			fmt.Println("---------------------")

			for i, item := range history {
				fmt.Print("log" + fmt.Sprint(i+1) + ": ")
				fmt.Println(item)
			}
			goto initial

		}

		if tokens == 6 {
			fmt.Println("shutting down the console!")
			os.Exit(0)
		}
		return
	}

}

func main() {
	powerHouse()
}
