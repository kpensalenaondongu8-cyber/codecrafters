package main

import (
	"fmt"
	"strconv"
	"strings"
)

func BinToDec(num string) int64 {
	v, err := strconv.ParseInt(num, 2, 64)
	if err != nil {
		fmt.Println("Invalid binary number")
		return 0
	}
	return v
}

func HexToDec(num string) int64 {
	v, err := strconv.ParseInt(num, 16, 64)
	if err != nil {
		fmt.Println("Invalid hex number")
		return 0
	}
	return v
}

func DecToBinHex(num int64) (string, string) {
	bin := strconv.FormatInt(num, 2)
	hex := strings.ToUpper(strconv.FormatInt(num, 16))
	return bin, hex
}

func baseConverter() {
	for {
		var convert string
		fmt.Println("\nChoose conversion:")
		fmt.Println("1. BinToDec")
		fmt.Println("2. HexToDec")
		fmt.Println("3. DecToBinHex")
		fmt.Println("4. quit")

		fmt.Print(">> ")
		fmt.Scanln(&convert)

		switch convert {

		case "BinToDec":
			var num string
			fmt.Print("Enter binary number: ")
			fmt.Scanln(&num)
			fmt.Println("Decimal:", BinToDec(num))
			history = append(history, convert+" "+num)

		case "HexToDec":
			var num string
			fmt.Print("Enter hex number: ")
			fmt.Scanln(&num)
			fmt.Println("Decimal:", HexToDec(num))
			history = append(history, convert+" "+num)

		case "DecToBinHex":
			var num int64
			fmt.Print("Enter decimal number: ")
			fmt.Scanln(&num)
			bin, hex := DecToBinHex(num)
			fmt.Println("Binary:", bin)
			fmt.Println("Hex   :", hex)
			history = append(history, convert+" "+fmt.Sprint(num))

		case "quit":
			fmt.Println("Okay Bye-bye")
			powerHouse()
			return

		default:
			fmt.Println("Invalid Input")
		}
	}
}
