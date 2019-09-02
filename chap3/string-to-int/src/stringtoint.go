/*
** AUTHOR	: Robert James Patterson
** DATE		: 08/28/2019
** SYNOPSIS	: We are going to use strconv.ParseInt and it returns two variables: the first one is
** the actual integer that we are expecting and the other is the return variable that arises if any
** error occurs during conversion.
 */
package main

import (
	"fmt"
	"strconv"
)

func main() {

	num := "104"
	valueInt, err := strconv.ParseInt(num, 10, 32)

	if err != nil {
		fmt.Printf("Some bad shit happend!! : %s \n", err)
	} else {
		fmt.Printf("Success! The sting value of 'num' has been converted to the int value: %d", valueInt)
	}
}
