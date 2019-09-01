/*
** AUTHOR	:	Robert James Patterson
** DATE		:	08/28
** SYNOPSIS	:	Type conversion using the 'strconv' pakage from the GO standard library. We are going
** to have a variable name isNew, and this is going to be a string value and is a true value. We're
** going to use a package called strconv, which has ParseBool. It returns two things: one is the
** Boolean value and the other is an error.
 */
package main

import (
	"fmt"
	"strconv"
)

func main() {

	isNew := "true"
	isNewBool, err := strconv.ParseBool(isNew)

	if err != nil {
		fmt.Println("failed")
	} else {
		if isNewBool {
			fmt.Printf("The value of \"IsNew\" : %s \n", strconv.FormatBool(isNewBool))
		} else {
			fmt.Println("not new")
		}
	}
}
