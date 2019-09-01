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

	// print with no trimming
	greeting := "\t Howdy, Bitches!! "
	fmt.Printf("%d %s\n", len(greeting), greeting)

	// trim spaces and print
	trimmed := strings.TrimSpace(greeting)
	fmt.Printf("%d %s\n", len(trimmed), trimmed)

	// slicing strings to extract substrings
	planets := "Venus and Mars"
	helloVenus := planets[0:5]
	helloMars := planets[10:14]
	fmt.Printf("Hello %s\n", helloVenus)
	fmt.Printf("Hello %s\n", helloMars)
}
