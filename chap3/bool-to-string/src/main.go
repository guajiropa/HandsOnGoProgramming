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

	isNew := true

	// This example uses the FormatBool function of the strconv package.
	isNewStr := strconv.FormatBool(isNew)
	message := "The purchased item is " + isNewStr

	// Print to the output device.
	fmt.Println("===========================")
	fmt.Printf("%s \n", message)
	fmt.Println("===========================")

	// This example uses the format 'verbs' to format the output.
	message = "The purchased item is "

	// Print to the output device.
	fmt.Println("===========================")
	fmt.Printf("%s%v \n", message, isNew)
	fmt.Println("===========================")
}
