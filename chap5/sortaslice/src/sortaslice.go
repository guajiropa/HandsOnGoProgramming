/*
** AUTHOR	: Robert James Patterson
** DATE		: 09/01/2019
** SYNOPSIS	: The other approach is to use sorting, where we can sort the slice first and then
** search for a particular item afterwards. In order to do that, Go provides a sort package. To
** be able to sort a slice, the slice needs to implement various methods that the sort package
** needs. The sort package provides a type called sort.stringslice, and what we can do is convert
** our stringslice to the StringSlice type provided by sort. Here, sortedList is not sorted, so
** we have to sort it explicitly.
 */
package main

import (
	"fmt"
	"sort"
)

func main() {

	// Setup our slice(array) of strings.
	strSlice := []string{"Sandy", "Provo", "St. George", "Salt Lake City", "Draper", "South Jordan", "Murray"}

	// Now iterate thru them looking for the value we desire.
	for idx, v := range strSlice {
		if v == "Sandy" {
			fmt.Println("============================================================")
			fmt.Printf("We found the string 'Sandy' at element %v of the slice.\n", idx)
			fmt.Println("============================================================")
			fmt.Printf("And our slice looks like this: %v\n", strSlice)
		}
	}
	sortedList := sort.StringSlice(strSlice)
	sortedList.Sort()

	// Now you can just use the Search fuction of sortedList to find your value.
	idx := sortedList.Search("Sandy")
	fmt.Println("============================================================")
	fmt.Printf("We found the string 'Sandy' at element %v of the sortedslice.\n", idx)
	fmt.Println("============================================================")
	fmt.Printf("And now our slice looks like this: %v\n", sortedList)

}
