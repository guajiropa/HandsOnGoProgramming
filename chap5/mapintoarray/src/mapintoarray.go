/*
** AUTHOR	: Robert James Patterson
** DATE		: 09/04/2019
** SYNOPSIS	: we are going to see how to convert a map into an array of keys and values. Let's
** imagine a variable named nameAges that has map, as shown in the following code block, and we map
** string values to integer values. There are names and ages too.
**
** We need to add a new struct named NameAge and it will have Name as a string and Age as integer.
** We are going to iterate over our nameAges map now. We'll be using a for loop and when you use a
** range operator on a map type, it returns two things, a key and the value.
 */
package main

import (
	"fmt"
)

type NameAge struct {
	// Struct to hold the Name & Age
	Name string
	Age  int
}

func main() {

	// Slice to hold the data from the struct
	var nameAgeSlice []NameAge

	nameAges := map[string]int{
		"Michelle": 30,
		"Jasmin":   25,
		"Savana":   23,
		"Vickey":   29,
		"Jessica":  32,
	}

	for key, value := range nameAges {
		nameAgeSlice = append(nameAgeSlice, NameAge{key, value})
	}

	fmt.Println(nameAgeSlice)
}
