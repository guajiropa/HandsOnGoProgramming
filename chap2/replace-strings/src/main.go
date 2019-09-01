/*
** AUTHOR	: Robert James Patterson
** DATE		: 08/21/2019
** SYNOPSIS	: Working thru the 'Hands-On Go Programming'
 */
package main

import (
	"fmt"
	"strings"
)

func main() {
	helloWorld := "Hello, World. How have you been World? \"I\" have been good, thanks for asking World."
	bookTitle := "live my own life"

	// using the 'replace' function from the 'strings' package. The minus one tells the function
	// to replace all instances it finds of the string 'World'.
	helloMars := strings.Replace(helloWorld, "World", "Mars", -1)
	fmt.Printf("%s\n%s\n", helloWorld, helloMars)
	fmt.Printf("Have you read the book:\t%s", strings.Title(bookTitle))
}
