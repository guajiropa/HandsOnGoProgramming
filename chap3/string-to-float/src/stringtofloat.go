/*
** AUTHOR	: Robert James Patterson
** DATE		: 08/29/2019
** SYNOPSIS	: We are going to use strconv.ParseFloat and it returns two variables: the first one is
** the actual float that we are expecting and the other is the return variable that arises if any
** error occurs during conversion.
 */
package main

import (
	"fmt"
	"strconv"
)

func main() {

	numstr := "247.015"
	valFloat, err := strconv.ParseFloat(numstr, 64)

	if err != nil {
		fmt.Printf("Something went wrong! %s \n", err)
	} else {
		fmt.Printf("The string was converted to the floating point value : %f ", valFloat)
	}
}
