/*
** AUTHOR	:	Robert James Patterson
** DATE		:	08/25
** SYNOPSIS	:	Type conversion using the 'strconv' pakage from the GO standard library.
 */
package main

import (
	"fmt"
	"strconv"
)

func main() {

	num := 100
	num64 := int64(num)

	numStr := strconv.FormatInt(num64, 10)
	binStr := strconv.FormatInt(num64, 2)

	// We can use Itoa instead to convert any integer to ASCII
	itoaStr := strconv.Itoa(num)

	floatnum := 856.98654

	// Use '-1' if you do not know the number of percision points
	floatnumStr := strconv.FormatFloat(floatnum, 'f', 2, 64)

	fmt.Println("Here is the float")
	fmt.Println(floatnumStr)
	fmt.Println("Now, we use 'FormatInt' with an Int64")
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
