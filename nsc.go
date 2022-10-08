package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/inancgumus/screen"
)

var number int64
var ns string

func main() {

	fmt.Println("A numeral system converter")
	converter()

}

func converter() {
	for {
		fmt.Print("Please enter a number: ")
		fmt.Scan(&number)
		fmt.Println("To what numeral system do you want to convert? (binary - B, octal - O, hexadecimal - H; quit - q)")
		fmt.Scan(&ns)
		ns = strings.ToLower(ns)
		if ns == "q" {
			os.Exit(0)
		}
		screen.Clear()
		screen.MoveTopLeft()
		if ns == "binary" || ns == "b" {
			number := strconv.FormatInt(number, 2)
			fmt.Println("-", number)
		} else if ns == "octal" || ns == "o" {
			number := strconv.FormatInt(number, 8)
			fmt.Println("-", number)
		} else if ns == "hexadecimal" || ns == "h" {
			number := strconv.FormatInt(number, 16)
			fmt.Println("-", strings.ToUpper(number))
		}
		fmt.Println("")
	}

}
