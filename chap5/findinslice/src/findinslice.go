/*
** AUTHOR	: Robert James Patterson
** DATE		: 09/01/2019
** SYNOPSIS	: We are going to learn how to find an element from an array or a slice. There are many
** ways you can do this, but we're going to cover two of them in this chapter. Imagine that we have
** a variable that holds a slice of strings. The first way of searching for a particular string in
** this slice will be using a for loop
 */
package main

import (
	"fmt"
)

func main() {

	// Setup our slice(array) of strings.
	strSlice := []string{"Sandy", "Provo", "St. George", "Salt Lake City", "Draper", "South Jordan", "Murray"}

	// Now iterate thru them looking for the value we desire.
	for idx, v := range strSlice {
		if v == "Sandy" {
			fmt.Println("========================================================")
			fmt.Printf("We found the string 'Sandy' at element %v of the slice.\n", idx)
			fmt.Println("========================================================")
		}
	}
}
