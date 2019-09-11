/*
** AUTHOR	: Robert James Patterson
** DATE		: 09/01/2019
** SYNOPSIS	: We are going to see how to check whether a key exists in a given map. So we have a map,
** nameAges, which basically maps names to ages.
 */
package main

import (
	"fmt"
)

func main() {

	nameAges := map[string]int{
		"Tarik":   32,
		"Michael": 30,
		"Jon":     25,
		"Jessica": 28,
	}

	if _, exists := nameAges["Jessica"]; exists {
		fmt.Printf("Jessica is %d years old.\n", nameAges["Jessica"])
	} else {
		fmt.Println("The key \"Jessica\" does not exist!")
	}
}
