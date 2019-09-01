/*
** AUTHOR	: Robert James Patterson
** DATE		: 09/01/2019
** SYNOPSIS	: n order for Go to understand this, you need to provide a layout attribute. A layout
** attribute basically describes how your string DateTime looks; it starts with a year, followed by
** a month, and day, and then the time. As always, the time package provides us with various utility
** functions that we can use to play with the date and time. The Parse method returns two things, one
** is the parse date and the other one is an error. If anything happens during parsing, an error can
** be thrown and we can check the error and see what went wrong, otherwise we will just output the
** string representation of the time now and the time we parsed.
 */
package main

import (
	"fmt"
	"time"
)

func main() {

	str := "2018-08-08T11:45:26.371Z"
	layout := "2006-01-02T15:04:05.000Z"

	t, err := time.Parse(layout, str)

	if err != nil {
		fmt.Printf("Something went wrong! ERROR: %v \n", err)
	} else {
		fmt.Printf("The string 'str' was converted to the DateTime value of : %s\n", t.String())
	}
}
