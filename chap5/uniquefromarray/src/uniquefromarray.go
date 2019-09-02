/*
** AUTHOR	: Robert James Patterson
** DATE		: 09/01/2019
** SYNOPSIS	: We are going to learn how to extract unique elements from a list. To begin, let's
** imagine we have a slice, containing duplicate elements in it. Now, let's say that we want to
** extract the unique elements. Since we don't have a built-in construct in Go, we're going to
** make our own function that will do the extraction. So, we have the uniqueIntSlice function
** and it accepts intSlice or intarray. Our unique function will accept intSlice and it will
** return another slice. So, the idea of this function is to track the duplicate elements in a
** separate list and if an element reappears in our given list, then we're not going to add
** that element to our new list.
 */
package main

import (
	"fmt"
)

func main() {

	// First, here is the slice with repetative values in it.
	intSlice := []int{1, 5, 5, 5, 5, 7, 8, 6, 6, 6}

	// Now let's us our function 'unique()' to get our unique values from intSlice.
	uniqueIntSlice := unique(intSlice)

	// Now display
	fmt.Println("================================================")
	fmt.Printf("The variable 'intSlice' contains the values %v\n", intSlice)
	fmt.Println("================================================")
	fmt.Printf("The unique values are %v\n", uniqueIntSlice)
	fmt.Println("================================================")
}

func unique(intSlice []int) []int {

	/*
	** To initialize a map, use the built in make function. The make function allocates and
	** initializes a hash map data structure and returns a map value that points to it. The
	** specifics of that data structure are an implementation detail of the runtime and are
	** not specified by the language itself.
	 */
	keys := make(map[int]bool)

	// Create an empty slice to hold the unique values.
	uniqueElements := []int{}

	/*
	** Loop thru the intSlice list to iterate over each element and add it to our new list if
	** it is not a duplicate. We basically get our value by passing an entry; if the value is
	** false, then we add that entry to our key or our map and set its value to true so that we
	** can see whether this element already appears in our list. We also have a built-in  append
	** function that accepts a slice and also appends the entry to the end of our slice, which
	** returns another slice.
	 */
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			uniqueElements = append(uniqueElements, entry)
		}
	}
	return uniqueElements
}
