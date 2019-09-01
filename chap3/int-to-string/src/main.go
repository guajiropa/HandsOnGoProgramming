/*
** AUTHOR	:	Robert James Patterson
** DATE		:	08/24
** SYNOPSIS	:	Type conversion using the 'strconv' pakage from the GO standard library.
 */
package main

import (
	"fmt"
	"strconv"
)

func main() {

	// We need to explicitly state that the type is of Int64 when using 'FormatInt'.
	num64 := int64(100)
	num := 100

	// Send the variable to convert and and the numerical base to convert to.
	numStr := strconv.FormatInt(num64, 10)
	binStr := strconv.FormatInt(num64, 2)

	// We can use Itoa instead to convert any integer to ASCII
	itoaStr := strconv.Itoa(num)

	// Dispaly the output to the screen.
	fmt.Println("First, we use 'FormatInt' with an Int64")
	fmt.Println("=========")
	fmt.Println(numStr)
	fmt.Println("=========")
	fmt.Println("Next we use Itoa instead wich will convert any integer to ASCII")
	fmt.Println("=========")
	fmt.Println(itoaStr)
	fmt.Println("=========")
	fmt.Println("When using 'FormatInt', we can convert numerical base on the int.")
	fmt.Printf("Which, in binary is %s", binStr)

}
