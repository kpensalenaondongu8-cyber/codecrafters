package main

import (
	"fmt"
	"strconv"
	"strings"
)

func calculator() {
	var num1 string
	var operator string
	var num2 string

	for {
		fmt.Println("Choose operation: Add, Sub, Mult, Div, quit, help")
		fmt.Scan(&operator)
		operator = strings.ToLower(operator)
		if operator == "help" {
			fmt.Println(" ")
			fmt.Println("GUIDELINES ARE\nAdd- Addition\n Sub- Substraction\n  Mult- Multiplication\n Div- Division\n quit- Exit program")
			continue
		}

		if operator == "quit" {
			fmt.Println("GoodBye And Thanks for your service")
			powerHouse()
			return

		}
		if operator != "add" &&
			operator != "sub" &&
			operator != "mult" &&
			operator != "div" {
			fmt.Println("Thrash")
		}
		fmt.Scan(&num1)

		fmt.Scan(&num2)
		num01, err := strconv.ParseFloat(num1, 64)
		if err != nil {
			fmt.Println("Error: Using Letters in place of Digits")
		}
		num02, err := strconv.ParseFloat(num2, 64)
		if err != nil {
			fmt.Println("Error: Using Letters in place of Digits")
		}

		if num02 == 0 && operator == "div" {
			fmt.Println("Error: try again can't divide by zero")
			continue
		} else if operator == "add" {
			fmt.Println("Result:", num01+num02)
			history = append(history, operator+" "+fmt.Sprint(num01)+" "+fmt.Sprint(num02))
		} else if operator == "sub" {
			fmt.Println("Result:", num01-num02)
			history = append(history, operator+" "+fmt.Sprint(num01)+" "+fmt.Sprint(num02))
		} else if operator == "mult" {
			fmt.Println("Result:", num01*num02)
			history = append(history, operator+" "+fmt.Sprint(num01)+" "+fmt.Sprint(num02))
		} else if operator == "div" {
			fmt.Println("Result:", num01/num02)
			history = append(history, operator+" "+fmt.Sprint(num01)+" "+fmt.Sprint(num02))
		}

	}
}
