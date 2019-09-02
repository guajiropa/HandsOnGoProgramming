/*
** AUTHOR	: Robert James Patterson
** DATE		: 09/02/2019
** SYNOPSIS	: We are going to learn how to iterate over an array. Iterating over an array is one of
** the most fundamental and common operations in Go programming. Sometimes, you won't need the index.
** In that case, you can just ignore it by using underscore (_). This will mean that you're only
** interested in the value.
 */
package main

import (
	"fmt"
)

func main() {

	numbers := []int{1, 5, 3, 6, 2, 10, 8}

	fmt.Println("=================================================")
	for index, value := range numbers {
		fmt.Printf("Index: %v has the value of : %v\n", index, value)
	}
	fmt.Println("=================================================")
}
