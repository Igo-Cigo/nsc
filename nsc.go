//TODO Add a save function and more numeral systems

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
		fmt.Println("To what numeral system do you want to convert? (decimal - D, binary - B, octal - O, hexadecimal - H; quit - q)")
		fmt.Scan(&ns)
		ns = strings.ToLower(ns)
		if ns == "q" {
			os.Exit(0)
		}
		screen.Clear()                   // Clears the console
		screen.MoveTopLeft()             // Exit causes the current program to exit with the given status code.
		if ns == "binary" || ns == "b" { //! Spaghetti code
			number := strconv.FormatInt(number, 2) // FormatInt returns the string representation of i in the given base, for 2 <= base <= 36. The result uses the lower-case letters 'a' to 'z' for digit values >= 10.
			fmt.Println("-", number)
		} else if ns == "octal" || ns == "o" {
			number := strconv.FormatInt(number, 8)
			fmt.Println("-", number)
		} else if ns == "hexadecimal" || ns == "h" {
			number := strconv.FormatInt(number, 16)
			fmt.Println("-", strings.ToUpper(number))
		} else if ns == "decimal" || ns == "d" { // very useful, no flaws whatsoever
			number := strconv.FormatInt(number, 10)
			fmt.Println("-", number)
		} else {
			os.Exit(0)
		}
	}
}
