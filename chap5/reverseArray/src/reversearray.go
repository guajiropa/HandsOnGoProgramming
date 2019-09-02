/*
** AUTHOR	: Robert James Patterson
** DATE		: 09/02/2019
** SYNOPSIS	: we are going to learn how to reverse sort an array. We'll have a variable that holds a
** slice of numbers. Since you are now familiar with the sort package in Go, you'll know that the
** sort package provides a lot of functionality that we can use to sort arrays and slices. If you
** look at the sort package, you'll see many types and functions. Now, we need the sort function and
** it accepts an interface, and this interface is defined in the sort package; therefore, we can call
** it the Sort interface. We are going to convert our slice of numbers into an interface.
 */
package main

import (
	"fmt"
	"sort"
)

func main() {

	numbers := []int{1, 5, 3, 6, 2, 10, 8}
	tobeSorted := sort.IntSlice(numbers)
	sort.Sort(tobeSorted)

	fmt.Println("================================================")
	fmt.Printf("We started out with an sorted list :\n %v\n", numbers)
	fmt.Println("================================================")

	sort.Sort(sort.Reverse(tobeSorted))

	fmt.Printf("After reversing we now have this : \n %v\n", tobeSorted)
	fmt.Println("================================================")
}
