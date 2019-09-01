/*
** AUTHOR	: Robert James Patterson
** DATE		: 08/29/2019
** SYNOPSIS	: Imagine that you have a helloWorldByte array; currently, it is a literal byte array,
** but you can derive it from any stream, such as a network or a file. We also have the string
** construct, which makes it really easy to convert a byte array into a string representation of
** it. We are going to use fmt.Println for the string representation of this helloWorldByte and
** run the code.
 */
package main

import (
	"fmt"
)

func main() {

	// Holds 'Howdy, Bitches!' as a byte array
	howdyBitchesByte := []byte{72, 111, 119, 100, 121, 44, 32, 66, 105, 116, 99, 104, 101, 115, 33}
	// Holds 'Howdy, Bitches!' as a string
	howdyBitchesString := "Howdy, Bitches!"

	// Covert a byte array to a string and display
	fmt.Println(string(howdyBitchesByte))
	// Convert a string to a byte array and display
	fmt.Println([]byte(howdyBitchesString))

}
